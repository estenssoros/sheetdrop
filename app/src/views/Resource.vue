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
    <b-breadcrumb>
      <b-breadcrumb-item>home</b-breadcrumb-item>
      <b-breadcrumb-item :href="'/orgs/' + resource.organization.id">{{
        resource.organization.name
      }}</b-breadcrumb-item>
      <b-breadcrumb-item active>{{ resource.name }}</b-breadcrumb-item>
    </b-breadcrumb>
    <b-row>
      <b-col>
        <h1>Schemas</h1>
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
      :items="schemas"
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
            :to="'/schema/' + data.item.id"
            variant="outline-primary"
          >
            <b-icon icon="arrow-right-square"></b-icon>
          </b-button>
          <b-button
            size="sm"
            v-on:click="deleteItem(data.item)"
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
      resource: {
        id: 0,
        name: "",
        organization: {
          id: 0,
          name: "",
        },
      },
      schemas: null,
      fields: [
        "id",
        "createdAt",
        "name",
        "startRow",
        "startColumn",
        "sourceType",
        "sourceURI",
        { key: "actions", label: "Actions" },
      ],
      routeParam: this.$route.params.id,
      file: null,
      name: "",
      nameState: null,
    };
  },
  computed: {
    formTitle() {
      return this.editing ? "Edit Schema" : "New Schema";
    },
  },
  methods: {
    closeModal() {
      this.name = "";
      this.nameState = null;
    },
    openModal() {
      if (this.editing) {
        return;
      }
      this.closeModal();
    },
    handleOk(bvModalEvt) {
      // Prevent modal from closing
      bvModalEvt.preventDefault();
      // Trigger submit handler
      this.handleSubmit();
    },
    async handleSubmit() {
      await this.$apollo.mutate({
        mutation: gql`
          mutation($resourceID: ID!, $schemaName: String!, $file: Upload!) {
            createSchemaWithFile(
              resourceID: $resourceID
              name: $schemaName
              file: $file
            ) {
              id
            }
          }
        `,
        variables: {
          schemaName: this.name,
          file: this.file,
          resourceID: this.resource.id,
        },
        context: {
          hasUpload: true,
        },
      });
      this.$apollo.queries.schemas.refetch();
      this.$bvModal.hide("modal-prevent-closing");
    },
  },
  apollo: {
    resource: {
      query: gql`
        query Resource($id: ID!) {
          resource(id: $id) {
            id
            name
            organization {
              id
              name
            }
          }
        }
      `,
      variables() {
        return {
          id: this.routeParam,
        };
      },
    },
    schemas: {
      query: gql`
        query Schemas($resourceID: ID!) {
          schemas(resourceID: $resourceID) {
            id
            createdAt
            name
            startRow
            startColumn
            sourceType
            sourceURI
          }
        }
      `,
      variables() {
        return {
          resourceID: this.routeParam,
        };
      },
    },
  },
};
</script>