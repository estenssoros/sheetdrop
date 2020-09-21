<template>
  <div>
    <b-modal
      id="modal-prevent-closing"
      ref="modal"
      :title="formTitle"
      @show="openModal"
      @hidden="resetModal"
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
      </form>
    </b-modal>
    <b-jumbotron header="APIs yo!"></b-jumbotron>
    <b-container>
      <b-button-toolbar class="float-right mb-1">
        <b-button v-b-modal.modal-prevent-closing variant="success">
          <b-icon icon="plus"></b-icon>Add New
        </b-button>
      </b-button-toolbar>
      <b-table hover responsive bordered :items="apis" :fields="fields" head-variant="light">
        <template v-slot:cell(actions)="data">
          <b-button-group class="mr-1">
            <b-button size="sm" v-on:click="editItem(data.item)" variant="outline-success">
              <b-icon icon="pencil-square"></b-icon>
            </b-button>
            <b-button size="sm" :to="'/api/'+data.item.ID" variant="outline-primary">
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

export default {
  data() {
    return {
      apis: [],
      fields: [
        "ID",
        "Name",
        { key: "actions", label: "Actions" },
      ],
      name: "",
      nameState: null,
      editingAPI: false,
      editedItem: null,
    };
  },
  created() {
    this.apiQuery();
  },
  computed: {
    formTitle() {
      return this.editingAPI ? "Edit API" : "New API";
    },
  },
  methods: {
    apiQuery() {
      client
        .getAPIs()
        .then((response) => {
          this.apis = response;
        })
        .catch((error) => {
          this.toastError(error);
        });
    },
    checkFormValidity() {
      const valid = this.$refs.form.checkValidity();
      this.nameState = valid;
      return valid;
    },
    openModal() {
      if (this.editingAPI) {
        return;
      }
      this.resetModal();
    },
    resetModal() {
      this.name = "";
      this.nameState = null;
    },
    handleOk(bvModalEvt) {
      bvModalEvt.preventDefault();
      this.handleSubmit();
    },
    handleSubmit() {
      if (!this.checkFormValidity()) {
        return;
      }
      if (this.editingAPI) {
        return this.updateAPI();
      }
      return this.createAPI();
    },
    updateAPI() {
      this.editedItem.Name = this.name;
      client
        .updateAPI(this.editedItem)
        .then((response) => {
          const index = this.apis.indexOf(this.editedItem);
          Object.assign(this.apis[index], response);
        })
        .catch((error) => {
          this.toastError(error);
        });
      this.$nextTick(() => {
        this.$bvModal.hide("modal-prevent-closing");
      });
    },
    createAPI() {
      client
        .createAPI({ Name: this.name })
        .then((response) => {
          this.apis.push(response);
        })
        .catch((error) => {
          this.toastError(error);
        });
      // Hide the modal manually
      this.$nextTick(() => {
        this.$bvModal.hide("modal-prevent-closing");
      });
    },
    deleteItem(item) {
      const index = this.apis.indexOf(item);
      if (confirm("Are you sure you want to delete this item?")) {
        client
          .deleteAPI(item)
          .then(() => {
            this.apis.splice(index, 1);
          })
          .catch((error) => {
            this.toastError(error);
          });
      }
    },

    toastError(error) {
      this.$bvToast.toast(error, {
        title: "API Error",
        variant: "danger",
        solid: true,
      });
    },
    editItem(item) {
      this.editingAPI = true;
      this.name = item.Name;
      this.editedItem = item;
      this.$bvModal.show("modal-prevent-closing");
    },
    followAPI(item) {
      this.setSelectedAPI(item);
      this.$router.replace("/api");
    },
  },
};
</script>