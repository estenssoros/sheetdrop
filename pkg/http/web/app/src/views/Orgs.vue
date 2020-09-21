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
      </form>
    </b-modal>
    <b-jumbotron header="Orgs yo!"></b-jumbotron>
    <b-container>
      <b-button-toolbar class="float-right mb-1">
        <b-button v-b-modal.modal-prevent-closing variant="success">
          <b-icon icon="plus"></b-icon>Add New
        </b-button>
      </b-button-toolbar>
      <b-table hover responsive bordered :items="items" :fields="fields" head-variant="light">
        <template v-slot:cell(actions)="data">
          <b-button-group class="mr-1">
            <b-button size="sm" v-on:click="editItem(data.item)" variant="outline-success">
              <b-icon icon="pencil-square"></b-icon>
            </b-button>
            <b-button size="sm" :to="'/orgs/'+data.item.ID" variant="outline-primary">
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
      items: [],
      fields: [
        "ID",
        "Name",
        "AccountLevel",
        "Members",
        { key: "actions", label: "Actions" },
      ],
      name: "",
      nameState: null,
      editing: false,
      editedItem: null,
    };
  },
  computed: {
    formTitle() {
      return this.editing ? "Edit Org" : "New Org";
    },
  },
  created() {
    this.orgsQuery();
  },
  methods: {
    closeModal() {
      this.name = "";
      this.nameState = null;
    },
    orgsQuery() {
      client
        .getOrgs()
        .then((response) => {
          this.items = response;
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
      if (this.editing) {
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
      if (this.editing) {
        return this.updateOrg();
      }
      return this.createOrg();
    },
    updateOrg() {
      this.editedItem.Name = this.name;
      client
        .updateOrg(this.editedItem)
        .then((response) => {
          const index = this.items.indexOf(this.editedItem);
          Object.assign(this.items[index], response);
        })
        .catch((error) => {
          this.toastError(error);
        });
      this.$nextTick(() => {
        this.$bvModal.hide("modal-prevent-closing");
      });
    },
    createOrg() {
      client
        .createOrg({ Name: this.name })
        .then((response) => {
          this.items.push(response);
        })
        .catch((error) => {
          console.log(error);
          this.toastError(error);
        });
      this.$nextTick(() => {
        this.$bvModal.hide("modal-prevent-closing");
      });
    },
    editItem(item) {
      this.editing = true;
      this.name = item.Name;
      this.editedItem = item;
      this.$bvModal.show("modal-prevent-closing");
    },
     deleteItem(item) {
      const index = this.items.indexOf(item);
      if (confirm("Are you sure you want to delete this item?")) {
        client
          .deleteOrg(item)
          .then(() => {
            this.items.splice(index, 1);
          })
          .catch((error) => {
            this.toastError(error);
          });
      }
    },
    toastError(error) {
      console.log(error)
      this.$bvToast.toast(error, {
        title: "API Error",
        variant: "danger",
        solid: true,
      });
    },
  },
};
</script>