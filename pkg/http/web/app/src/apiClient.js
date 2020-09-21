import Vue from 'vue';
import axios from 'axios';

const BASE_URI = 'http://localhost:1323/api';
const client = axios.create({
    baseURL: BASE_URI,
    json: true
});

const Client = {
    createAPI(api) {
        return this.post('/api', api);
    },

    deleteAPI(api) {
        return this.delete(`/api`, api);
    },

    updateAPI(api) {
        return this.patch(`/api`, api);
    },

    getAPIs() {
        return this.get('/apis');
    },

    getAPI(apiID) {
        return this.get(`/api/${apiID}`);
    },
    getOrgs() {
        return this.get('/orgs');
    },
    createOrg(org) {
        return this.post('/org', org)
    },
    updateOrg(org) {
        return this.patch('/org', org)
    },
    deleteOrg(org) {
        return this.delete('/org', org)
    },

    getSchemas(apiID) {
        return this.get(`/schema/${apiID}`);
    },
    updateSchema(schema) {
        return this.patch(`/schema`, schema)
    },
    deleteSchema(schema) {
        return this.delete("/schema", schema)
    },

    get(resource, data) {
        return this.perform('get', resource, data)
    },

    post(resource, data) {
        return this.perform('post', resource, data)
    },
    patch(resource, data) {
        return this.perform('patch', resource, data)
    },
    delete(resource, data) {
        return this.perform('delete', resource, data)
    },

    async perform(method, resource, data) {
        let accessToken = await Vue.prototype.$auth.accessToken
        return client({
            method,
            url: resource,
            data,
            headers: {
                Authorization: `Bearer ${accessToken}`
            }
        }).then(req => {
            return req.data
        })
    }
}
export default Client;