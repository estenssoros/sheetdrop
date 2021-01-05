<template>
  <div>
    <b-breadcrumb>
      <b-breadcrumb-item>home</b-breadcrumb-item>
      <b-breadcrumb-item :href="'/orgs/' + schema.resource.organization.id">{{
        schema.resource.organization.name
      }}</b-breadcrumb-item>
      <b-breadcrumb-item :href="'/resource/' + schema.resource.id">{{
        schema.resource.name
      }}</b-breadcrumb-item>
      <b-breadcrumb-item active>{{ schema.name }}</b-breadcrumb-item>
    </b-breadcrumb>
    <h1>Headers</h1>
    <b-table
      hover
      responsive
      bordered
      small
      :items="headers"
      :fields="fields"
      head-variant="light"
    >
      <template v-slot:cell(isID)="data">
        <b-form-checkbox v-model="data.item.isID" :disabled="true" />
      </template>
      <template v-slot:cell(actions)="data">
        <b-button-group class="mr-1">
          <b-button
            v-if="!data.item.isID"
            size="sm"
            v-on:click="setID(data.item, true)"
            variant="outline-success"
          >
            set id
          </b-button>
          <b-button
            v-if="data.item.isID"
            size="sm"
            v-on:click="setID(data.item, false)"
            variant="outline-danger"
          >
            remove id
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
      schema: {
        id: 0,
        name: "",
        resource: {
          id: 0,
          name: "",
          organization: {
            id: 0,
            name: "",
          },
        },
      },
      headers: null,
      routeParam: this.$route.params.id,
      fields: [
        "id",
        "name",
        "index",
        "dataType",
        "isID",
        { key: "actions", label: "Actions" },
      ],
    };
  },
  methods: {
    async setID(header, isID) {
      await this.$apollo.mutate({
        mutation: gql`
          mutation($headerID: ID!, $isID: Boolean!) {
            setHeaderID(id: $headerID, isID: $isID) {
              id
            }
          }
        `,
        variables: {
          headerID: header.id,
          isID: isID,
        },
      });
      this.$apollo.queries.headers.refetch();
    },
  },
  apollo: {
    schema: {
      query: gql`
        query Schema($id: ID!) {
          schema(id: $id) {
            id
            name
            resource {
              id
              name
              organization {
                id
                name
              }
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
    headers: {
      query: gql`
        query Headers($schemaID: ID!) {
          headers(schemaID: $schemaID) {
            id
            name
            index
            dataType
            isID
          }
        }
      `,
      variables() {
        return {
          schemaID: this.routeParam,
        };
      },
    },
  },
};
</script>