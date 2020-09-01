import Vue from 'vue';
import axios from 'axios';

const BASE_URI = 'http://localhost:1323/api';
const client = axios.create({
    baseURL: BASE_URI,
    headers: {
        "Content-Type": "multipart/form-data",
    }
});

const FormClient = {
    uploadSchemaFile(data) {
        return this.upload('patch', '/file-upload', data)
    },

    createNewSchema(data){
        return this.upload('post', '/file-upload', data)
    },

    async upload(method, resource, data) {
        let formData = new FormData();
        for (let [key, value] of Object.entries(data)) {
            formData.append(key, value)
        }
        let accessToken = await Vue.prototype.$auth.getAccessToken()
        return client({
            method,
            url: resource,
            data: formData,
            headers: {
                "Content-Type": "multipart/form-data",
                Authorization: `Bearer ${accessToken}`
            },
        }).then(req => {
            return req.data
        })
    }
}

export default FormClient