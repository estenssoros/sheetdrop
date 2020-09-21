<template>
  <div>
    <b-modal
      id="modal-prevent-closing"
      ref="modal"
      :title="formTitle"
      @show="openModal"
      @hidden="closeModal"
      @ok="handleOk"
    >
      <form ref="form" @submit.stop.prevent="handleSubmit">
        <b-form-group
          :state="nameState"
          label="Name"
          label-for="name-input"
          invalid-feedback="Name is required"
        >
          <b-form-input id="name-input" v-model="name" :state="nameState" required></b-form-input>
        </b-form-group>
        <b-form-group>
          <b-form-file
            v-model="file"
            :state="Boolean(file)"
            placeholder="Choose a file or drop it here..."
            drop-placeholder="Drop file here..."
            required
          ></b-form-file>
        </b-form-group>
      </form>
    </b-modal>
    <div>
      <b-sidebar
        id="sidebar-no-header"
        aria-labelledby="sidebar-no-header-title"
        no-header
        shadow
        backdrop
        right
        width="50%"
        lazy
        :visible="drawerVisible"
      >
        <template>
          <div class="p-3">
            <h4 id="sidebar-no-header-title">{{selectedSchema.Name}}</h4>
            <p>
              Cras mattis consectetur purus sit amet fermentum. Cras justo odio, dapibus ac facilisis
              in, egestas eget quam. Morbi leo risus, porta ac consectetur ac, vestibulum at eros.
            </p>
            <b-card header="Schema Information" class="mb-1">
              <b-list-group>
                <b-list-group-item>
                  Source URI:
                  <b-badge>{{selectedSchema.SourceURI}}</b-badge>
                </b-list-group-item>
                <b-list-group-item>
                  Source Type:
                  <b-badge>{{selectedSchema.SourceType}}</b-badge>
                </b-list-group-item>
                <b-list-group-item>
                  Start Column:
                  <b-badge>{{selectedSchema.StartColumn}}</b-badge>
                </b-list-group-item>
                <b-list-group-item>
                  Start Row:
                  <b-badge>{{selectedSchema.StartRow}}</b-badge>
                </b-list-group-item>
                <b-list-group-item>
                  AuthToken:
                  <b-form-input v-model="selectedSchema.AuthToken" disabled></b-form-input>
                </b-list-group-item>
              </b-list-group>
            </b-card>
            <b-card header="Data Headers" class="mb-1">
              <b-table
                :items="selectedSchema.Headers"
                :fields="schemaFields"
                hover
                small
                responsive
              >
                <template v-slot:cell(IsID)="data">
                  <b-form-checkbox v-model="data.item.IsID" v-on:change="updateID(data.item)" ></b-form-checkbox>
                </template>
              </b-table>
            </b-card>
            <b-button-toolbar>
              <b-button-group>
                <b-button variant="primary" block @click="hideDrawer()">Close</b-button>
              </b-button-group>
            </b-button-toolbar>
          </div>
        </template>
      </b-sidebar>
    </div>
    <b-jumbotron :header="api.Name + ' Schemas'"></b-jumbotron>
    <b-container>
      <b-button-toolbar class="float-right mb-1">
        <b-button v-b-modal.modal-prevent-closing variant="success">
          <b-icon icon="plus"></b-icon>Add New
        </b-button>
      </b-button-toolbar>
      <b-table hover responsive bordered :items="schemas" :fields="fields" head-variant="light">
        <template v-slot:cell(name)="data" v-on:click=" toggleName(data.item)">{{data.item.Name}}</template>
        <template v-slot:cell(actions)="data">
          <b-button-group class="mr-1">
            <b-button size="sm" v-on:click="editItem(data.item)" variant="outline-success">
              <b-icon icon="pencil-square"></b-icon>
            </b-button>
            <b-button size="sm" v-on:click="setSchema(data.item)" variant="outline-primary">
              <b-icon icon="arrow-right-square"></b-icon>
            </b-button>
            <b-button size="sm" v-on:click="deleteItem(data.item)" variant="outline-danger">
              <b-icon icon="trash" variant="danger"></b-icon>
            </b-button>
          </b-button-group>
        </template>
      </b-table>
    </b-container>
  </div>
</template>
<script>
import client from "../apiClient";
import formClient from "../formClient";

export default {
  data() {
    return {
      api: {},
      schemas: [],
      fields: [
        "ID",
        "Name",
        "SourceType",
        "SourceURI",
        "StartColumn",
        "StartRow",
        { key: "actions", label: "Actions" },
      ],
      name: "",
      nameState: null,
      file: null,
      editing: false,
      editedItem: {},
      selectedSchema: {},
      drawerVisible: false,
      schemaFields: ["Index", "Name", "DataType", "IsID"],
    };
  },
  computed: {
    formTitle() {
      return this.editing ? "Edit Schema" : "New Schema";
    },
  },
  created() {
    this.apiQuery(this.$route.params.id);
    this.schemaQuery(this.$route.params.id);
  },
  methods: {
    updateID(schema){
      console.log(schema)
    },
    hideDrawer() {
      this.drawerVisible = false;
    },
    setSchema(schema) {
      this.selectedSchema = schema;
      this.drawerVisible = true;
    },
    apiQuery(apiID) {
      client
        .getAPI(apiID)
        .then((response) => {
          this.api = response;
        })
        .catch((error) => {
          this.toastError(error);
        });
    },
    schemaQuery(apiID) {
      client
        .getSchemas(apiID)
        .then((response) => {
          this.schemas = response;
        })
        .catch((error) => {
          this.toastError(error);
        });
    },
    openModal() {
      if (this.editing) {
        return;
      }
      this.closeModal();
    },
    closeModal() {
      this.name = "";
      this.nameState = null;
    },
    handleOk(bvModalEvt) {
      // Prevent modal from closing
      bvModalEvt.preventDefault();
      // Trigger submit handler
      this.handleSubmit();
    },
    checkFormValidity() {
      const valid = this.$refs.form.checkValidity();
      this.nameState = valid;
      return valid;
    },
    handleSubmit() {
      // Exit when the form isn't valid
      if (!this.checkFormValidity()) {
        return;
      }
      if (this.editing) {
        this.editing = false;
        return this.updateSchema();
      }
      this.editing = false;
      this.createSchema();
    },
    updateSchema() {
      var index = this.schemas.indexOf(this.editedItem);
      var item = this.editedItem;
      if (this.file === undefined) {
        item.name = this.name;
        client
          .updateSchema(item)
          .then((response) => {
            Object.assign(this.schemas[index], response);
          })
          .catch((error) => {
            this.toastError(error);
          });
      } else {
        item.file = this.file;
        item.name = this.name;
        formClient
          .uploadSchemaFile(item)
          .then((response) => {
            Object.assign(this.schemas[index], response);
          })
          .catch((error) => {
            this.toastError(error);
          });
      }
      this.$nextTick(() => {
        this.$bvModal.hide("modal-prevent-closing");
      });
    },
    createSchema() {
      formClient
        .createNewSchema({
          id: 0,
          name: this.name,
          file: this.file,
          api_id: this.$route.params.id,
        })
        .then((response) => {
          this.schemas.push(response);
        })
        .catch((error) => {
          this.toastError(error);
        });
      this.$nextTick(() => {
        this.$bvModal.hide("modal-prevent-closing");
      });
    },
    deleteItem(item) {
      const index = this.schemas.indexOf(item);
      if (confirm("Are you sure you want to delete this item?")) {
        client
          .deleteSchema(item)
          .then(() => {
            this.schemas.splice(index, 1);
          })
          .catch((error) => {
            this.toastError(error);
          });
      }
    },
    editItem(item) {
      this.editing = true;
      this.name = item.Name;
      this.editedItem = item;
      this.$bvModal.show("modal-prevent-closing");
    },
    toastError(error) {
      this.$bvToast.toast(error, {
        title: "API Error",
        variant: "danger",
        solid: true,
      });
    },
  },
};
</script>