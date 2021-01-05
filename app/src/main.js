import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

import VueApollo from "vue-apollo";
import { ApolloClient } from "apollo-client";
// import { createHttpLink } from "apollo-link-http";
import { InMemoryCache } from "apollo-cache-inmemory";
import { createUploadLink } from "apollo-upload-client";
import { ApolloLink } from "apollo-link"

Vue.config.productionTip = false;

Vue.use(VueApollo);

const cache = new InMemoryCache();

const apolloClient = new ApolloClient({
  link: ApolloLink.from([
    // createHttpLink({
    //   uri: "http://localhost:1323/graph",
    // }),
    createUploadLink({
      uri: "http://localhost:1323/graph",
    }),
  ]),
  cache,
});

const apolloProvider = new VueApollo({
  defaultClient: apolloClient,
});

import { BootstrapVue, IconsPlugin } from "bootstrap-vue";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";

Vue.use(BootstrapVue);
Vue.use(IconsPlugin);

new Vue({
  router,
  store,
  apolloProvider,
  render: (h) => h(App),
}).$mount("#app");
