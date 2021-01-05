<template>
  <div>
    <b-navbar toggleable="lg" type="dark" variant="dark">
      <b-navbar-brand href="#">SheetDrop</b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
          <b-nav-item to="/">Home</b-nav-item>
          <b-nav-item to="/about">About</b-nav-item>
          <b-nav-item-dropdown text="Organizations">
            <b-dropdown-item
              v-for="org in organizations"
              :key="org.id"
              :to="'/orgs/' + org.id"
            >
              {{ org.name }}
            </b-dropdown-item>
            <b-dropdown-divider></b-dropdown-divider>
            <b-dropdown-item to="/orgs">
              <b-icon icon="plus"></b-icon>Add
            </b-dropdown-item>
          </b-nav-item-dropdown>
          <b-nav-item to="/apis">APIs</b-nav-item>
        </b-navbar-nav>

        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto">
          <b-button @click="login" size="sm">Login</b-button>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </div>
</template>

<script>
import gql from "graphql-tag";
export default {
  data() {
    return {
      organizations: null,
    };
  },
  apollo: {
    organizations: gql`
    query{
        organizations {
            id
            name
        }
    }`,
  },
  created() {},
  methods: {
    login() {},
    logout() {},
  },
};
</script>

<style lang="scss" scoped>
#nav a {
  color: white;
}
</style>