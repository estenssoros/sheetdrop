import Vue from 'vue';
import axios from 'axios';

const BASE_URI = 'http://localhost:1323/api';
const client = axios.create({
    baseURL: BASE_URI,
    json: true
});

const APIClient = {
    createAPI(api) {
        return this.perform('post', '/api', api);
    },

    deleteAPI(api) {
        return this.perform('delete', `/api`, api);
    },

    updateAPI(api) {
        return this.perform('patch', `/api`, api);
    },

    getAPIs() {
        return this.perform('get', '/api');
    },

    getAPI(api) {
        return this.perform('get', `/api/${api.id}`);
    },

    getSchemas(apiID){
        return this.perform('get',`/schema/${apiID}`);
    },
    updateSchema(schema){
        return this.perform('patch',`/schema`,schema)
    },

    async perform(method, resource, data) {
        let accessToken = await Vue.prototype.$auth.getAccessToken()
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

export default APIClient;