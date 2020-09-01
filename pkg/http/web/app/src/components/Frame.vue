<template>
  <div>
    <v-navigation-drawer v-model="drawer" app>
      <v-list dense>
        <v-list-item link>
          <v-list-item-action>
            <v-icon>mdi-home</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Home</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item link>
          <v-list-item-action>
            <v-icon>mdi-email</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>APIs</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar app color="indigo" dark>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>SheetDrop</v-toolbar-title>
      <v-spacer></v-spacer>
      <button @click.prevent="logout">Logout</button>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  props: {},
  data: () => ({
    drawer: null,
  }),
  methods: {
    async logout() {
      await this.$auth.logout();
      this.$router.push("/");
    },
  },
};
</script>