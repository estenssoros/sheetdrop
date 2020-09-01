<template>
  <v-app>
    <Frame />
    <v-main>
      <v-data-table
        :headers="apiHeaders"
        :items="apis"
        :single-select="true"
        item-key="ID"
        show-select
        dense
        v-on:item-selected="handleSelection"
        class="elevation-1"
      >
        <template v-slot:top>
          <v-toolbar flat color="white">
            <v-toolbar-title>My APIs</v-toolbar-title>
            <v-divider class="mx-4" inset vertical></v-divider>
            <v-spacer></v-spacer>
            <v-dialog v-model="dialog" max-width="500px">
              <template v-slot:activator="{ on, attrs }">
                <v-btn color="primary" dark class="mb-2" v-bind="attrs" v-on="on">New API</v-btn>
              </template>
              <v-card>
                <v-card-title>
                  <span class="headline">{{ formTitle }}</span>
                </v-card-title>

                <v-card-text>
                  <v-container>
                    <v-row>
                      <v-col cols="12" sm="6" md="4">
                        <v-text-field v-model="editedItem.Name" label="Name"></v-text-field>
                      </v-col>
                    </v-row>
                  </v-container>
                </v-card-text>

                <v-card-actions>
                  <v-spacer></v-spacer>
                  <v-btn color="blue darken-1" text @click="close">Cancel</v-btn>
                  <v-btn color="blue darken-1" text @click="save">Save</v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
          </v-toolbar>
        </template>
        <template v-slot:item.actions="{ item }">
          <v-icon small class="mr-2" @click="editItem(item)">mdi-pencil</v-icon>
          <v-icon small @click="deleteItem(item)">mdi-delete</v-icon>
        </template>
      </v-data-table>
      <template v-if="showSchemas">
        <v-divider />
        <Schemas :ApiID="selectedAPI.ID" :ApiName="selectedAPI.Name" />
      </template>
    </v-main>
  </v-app>
</template>

<script>
import Frame from "./Frame";
import APIClient from "../apiClient";
import Schemas from "./Schemas";
import "vuex";

export default {
  name: "Home",
  components: { Frame, Schemas },
  data() {
    return {
      apis: [],
      selectedAPI: {},
      schemas: [],
      showSchemas: false,
      dialog: false,
      apiHeaders: [
        { text: "id", value: "ID" },
        { text: "created at", value: "CreatedAt" },
        { text: "updated at", value: "UpdatedAt" },
        { text: "name", value: "Name" },
        { text: "Actions", value: "actions", sortable: false },
      ],
      editedIndex: -1,
      editedItem: {
        ID: -1,
        Name: "",
      },
      defaultItem: {
        ID: -1,
        Name: "",
      },
    };
  },

  computed: {
    formTitle() {
      return this.editedIndex === -1 ? "New API" : "Edit API";
    },
    schemaFormTitle() {
      return this.editedIndex === -1 ? "New Schema" : "Edit Schema";
    },
  },

  watch: {
    dialog(val) {
      val || this.close();
    },
  },

  created() {
    this.apisQuery();
  },

  methods: {
    apisQuery() {
      APIClient.getAPIs()
        .then((response) => {
          this.apis = response;
        })
        .catch((error) => {
          console.log(error);
        });
    },

    close() {
      this.dialog = false;
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem);
        this.editedIndex = -1;
      });
    },

    editItem(item) {
      this.editedIndex = this.apis.indexOf(item);
      this.editedItem = Object.assign({}, item);
      this.dialog = true;
    },

    deleteItem(item) {
      const index = this.apis.indexOf(item);
      if (confirm("Are you sure you want to delete this item?")) {
        APIClient.deleteAPI(item)
          .then(() => {
            this.apis.splice(index, 1);
          })
          .catch((error) => {
            console.log(error);
          });
      }
    },

    save() {
      var memIndex = this.editedIndex;
      if (memIndex > -1) {
        APIClient.updateAPI(this.editedItem)
          .then((response) => {
            Object.assign(this.apis[memIndex], response);
          })
          .catch((error) => {
            console.log(error);
          });
      } else {
        APIClient.createAPI(this.editedItem)
          .then((response) => {
            this.apis.push(response);
          })
          .catch((error) => {
            console.log(error);
          });
      }
      this.close();
    },

    handleSelection(item) {
      this.showSchemas = item.value;
      this.selectedAPI = item.item;
    },
  },
};
</script>

<style>
.v-tabs__content {
  padding-bottom: 2px;
}
</style>