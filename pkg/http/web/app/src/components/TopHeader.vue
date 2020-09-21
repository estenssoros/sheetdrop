<template>
  <div>
    <b-navbar toggleable="lg" type="dark" variant="dark">
      <b-navbar-brand href="#">SheetDrop</b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
          <b-nav-item to="/">Home</b-nav-item>
          <b-nav-item to="/about">About</b-nav-item>
          <div v-if="$auth.isAuthenticated()">
            <b-nav-item-dropdown text="Organizations" >
              <b-dropdown-item v-for="org in orgs" :key="org.ID" :to="'/orgs/'+org.ID">
                  {{org.Name}}
                </b-dropdown-item>
              <b-dropdown-divider></b-dropdown-divider>
              <b-dropdown-item to="/orgs">
                <b-icon icon="plus"></b-icon>Add
              </b-dropdown-item>
            </b-nav-item-dropdown>
          </div>
          <div v-if="$auth.isAuthenticated()">
            <b-nav-item to="/apis">APIs</b-nav-item>
          </div>
        </b-navbar-nav>

        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto">
          <!-- <b-nav-form>
            <b-form-input size="sm" class="mr-sm-2" placeholder="Search"></b-form-input>
            <b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
          </b-nav-form>-->
          <div v-if="$auth.isAuthenticated()">
            <b-nav-item-dropdown right>
              <!-- Using 'button-content' slot -->
              <template v-slot:button-content>
                <!-- <em style="color:white;">{{ $auth.user.name }}</em> -->
                <b-avatar :src="$auth.user.picture"></b-avatar>
              </template>
              <b-dropdown-item href="/profile">Profile</b-dropdown-item>
              <b-dropdown-item @click="logout">Sign Out</b-dropdown-item>
            </b-nav-item-dropdown>
          </div>
          <div v-if="!$auth.isAuthenticated()">
            <b-button @click="login" size="sm">Login</b-button>
          </div>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </div>
</template>

<script>
import client from "../apiClient";
export default {
  data(){
    return {
      orgs:[]
    }
  },
  created(){
    if (this.$auth.isAuthenticated()){
      this.getOrgs()
    }
  },
  methods: {
    login() {
      this.$auth.login();
    },
    logout() {
      this.$auth.logout({
        returnTo: window.location.origin,
      });
    },
    getOrgs() {
      client
        .getOrgs()
        .then((response) => {
          this.orgs = response;
        })
        .catch((error) => {
          this.toastError(error);
        });
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

<style lang="scss" scoped>
#nav a {
  color: white;
}
</style>