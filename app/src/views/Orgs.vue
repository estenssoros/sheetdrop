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
          <b-form-input
            id="name-input"
            v-model="name"
            :state="nameState"
            required
          ></b-form-input>
        </b-form-group>
      </form>
    </b-modal>
    <b-breadcrumb>
      <b-breadcrumb-item>home</b-breadcrumb-item>
      <b-breadcrumb-item active>{{ organization.name }}</b-breadcrumb-item>
    </b-breadcrumb>
    <b-row>
      <b-col>
        <h1>Resources</h1>
      </b-col>
      <b-col>
        <b-button-toolbar class="float-right mb-1">
          <b-button v-b-modal.modal-prevent-closing variant="success">
            <b-icon icon="plus"></b-icon>Add New
          </b-button>
        </b-button-toolbar>
      </b-col>
    </b-row>
    <b-table
      hover
      responsive
      bordered
      small
      :items="resources"
      :fields="fields"
      head-variant="light"
    >
      <template v-slot:cell(actions)="data">
        <b-button-group class="mr-1">
          <b-button
            size="sm"
            v-on:click="editItem(data.item)"
            variant="outline-success"
          >
            <b-icon icon="pencil-square"></b-icon>
          </b-button>
          <b-button
            size="sm"
            :to="'/resource/' + data.item.id"
            variant="outline-primary"
          >
            <b-icon icon="arrow-right-square"></b-icon>
          </b-button>
          <b-button
            size="sm"
            v-on:click="deleteResource(data.item)"
            variant="outline-danger"
          >
            <b-icon icon="trash" variant="danger"></b-icon>
          </b-button>
        </b-button-group>
      </template>
    </b-table>
  </div>
</template>
<script>
import gql from "graphql-tag";
export default {
  data() {
    return {
      resources: null,
      organization: {
        id: 0,
        name: "",
      },
      fields: [
        "id",
        "name",
        "createdAt",
        "schemaCount",
        { key: "actions", label: "Actions" },
      ],
      routeParam: this.$route.params.id,
      name: "",
      nameState: null,
      editing: false,
    };
  },
  computed: {
    formTitle() {
      return this.editing ? "Edit Org" : "New Org";
    },
  },
  methods: {
    closeModal() {
      this.name = "";
      this.nameState = null;
    },
    handleOk(bvModalEvt) {
      bvModalEvt.preventDefault();
      this.handleSubmit();
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
    checkFormValidity() {
      const valid = this.$refs.form.checkValidity();
      this.nameState = valid;
      return valid;
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
    async deleteResource(item) {
      // const index = this.items.indexOf(item);
      if (confirm("Are you sure you want to delete this item?")) {
        await this.$apollo.mutate({
          mutation: gql`
            mutation($resourceID: ID!) {
              deleteResource(id: $resourceID) {
                id
                name
              }
            }
          `,
          variables: {
            resourceID: item.id,
          },
        });
        this.$apollo.queries.resources.refetch()
        console.log(this.$apollo.queries)
      }
    },
    // updateOrg() {},
    async createOrg() {
      const results = await this.$apollo.mutate({
        mutation: gql`
          mutation($orgID: ID!, $resourceName: String!) {
            createResource(orgID: $orgID, resourceName: $resourceName) {
              id
              createdAt
              name
              schemaCount
            }
          }
        `,
        variables: {
          resourceName: this.name,
          orgID: this.organization.id,
        },
      });
      this.resources.push(results.data.createResource);
      this.$bvModal.hide("modal-prevent-closing");
    },
  },
  apollo: {
    resources: {
      query: gql`
        query Resources($organizationID: ID!) {
          resources(organizationID: $organizationID) {
            id
            createdAt
            name
            schemaCount
          }
        }
      `,
      variables() {
        return {
          organizationID: this.routeParam,
        };
      },
    },
    organization: {
      query: gql`
        query Organization($id: ID!) {
          organization(id: $id) {
            id
            name
          }
        }
      `,
      variables() {
        return {
          id: this.routeParam,
        };
      },
    },
  },
};
</script>