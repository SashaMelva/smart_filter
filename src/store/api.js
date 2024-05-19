import axios from "axios";

const instance = axios.create({
    withCredentials: true,
    baseURL: 'http://localhost:8000/'
});

export const userAPI = {
    getMe() {
        return instance.get("auth/me")
            .then(response => {
                return response.data;
            })
            .catch(err => {
                return null;
            });
    }
}
