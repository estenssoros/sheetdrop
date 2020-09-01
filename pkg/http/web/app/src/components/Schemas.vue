<template>
  <v-data-table :headers="headers" :items="schemas" item-key="ID" dense>
    <template v-slot:top>
      <v-toolbar flat color="white">
        <v-toolbar-title>{{ ApiName }} Schemas</v-toolbar-title>
        <v-divider class="mx-4" inset vertical></v-divider>
        <v-spacer></v-spacer>
        <v-dialog v-model="dialog" max-width="500px">
          <template v-slot:activator="{ on, attrs }">
            <v-btn color="secondary" dark class="mb-2" v-bind="attrs" v-on="on">New Schema</v-btn>
          </template>
          <v-card>
            <v-card-title>
              <span class="headline">{{ formTitle }}</span>
            </v-card-title>

            <v-card-text>
              <v-container>
                <v-row>
                  <v-col cols="12" sm="12" md="12">
                    <v-text-field v-model="editedItem.Name" label="Name"></v-text-field>
                  </v-col>
                </v-row>
                <v-col cols="12" sm="12" md="12">
                  <v-file-input
                    id="file"
                    ref="file"
                    small-chips
                    dense
                    label="Upload"
                    @change="selectFile"
                  ></v-file-input>
                </v-col>
                <v-row></v-row>
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
</template>
 
<script>
import APIClient from "../apiClient";
import FormClient from "../formClient";

export default {
  name: "Schemas",
  props: {
    ApiID: Number,
    ApiName: String,
  },
  data() {
    return {
      currentFile: undefined,
      schemas: [],
      dialog: false,
      headers: [
        { text: "id", value: "ID" },
        { text: "created at", value: "CreatedAt" },
        { text: "updated at", value: "UpdatedAt" },
        { text: "name", value: "Name" },
        { text: "source type", value: "SourceType" },
        { text: "source uri", value: "Source URI" },
        { text: "actions", value: "actions", sortable: false },
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
      return this.editedIndex === -1 ? "New Schema" : "Edit Schema";
    },
  },
  watch: {
    dialog(val) {
      val || this.close();
    },
  },
  created() {
    this.SchemaQuery();
  },
  methods: {
    SchemaQuery() {
      APIClient.getSchemas(this.$props.ApiID)
        .then((response) => {
          this.schemas = response;
          this.showSchemas = true;
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
    deleteItem(item){
      const index = this.schemas.indexOf(item);
      if (confirm("Are you sure you want to delete this item?")) {
        APIClient.deleteSchema(item)
          .then(() => {
            this.schemas.splice(index, 1);
          })
          .catch((error) => {
            console.log(error);
          });
      }
    },
    save() {
      var memIndex = this.editedIndex;
      if (memIndex > -1) {
        if (this.currentFile === undefined) {
          APIClient.updateSchema(this.editedItem)
            .then((response) => {
              Object.assign(this.schemas[memIndex], response);
            })
            .catch((error) => {
              console.log(error);
            });
          this.close();
          return;
        }
        var schema = this.editedItem;
        schema.file = this.currentFile;
        FormClient.uploadSchemaFile(schema)
          .then((response) => {
            Object.assign(this.schemas[memIndex], response);
          })
          .catch((error) => {
            console.log(error);
          });
      } else {
        var createSchema = this.editedItem;
        createSchema.file = this.currentFile;
        createSchema.api_id = this.ApiID;
        FormClient.createNewSchema(createSchema)
          .then((response) => {
            Object.assign(this.schemas[memIndex], response);
          })
          .catch((error) => {
            console.log(error);
          });
      }
      this.close();
    },

    editItem(item) {
      this.editedIndex = this.schemas.indexOf(item);
      this.editedItem = Object.assign({}, item);
      this.dialog = true;
    },

    selectFile(file) {
      this.currentFile = file;
    },
  },
};
</script>