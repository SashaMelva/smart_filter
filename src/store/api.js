import axios from "axios";

const instance = axios.create({
    withCredentials: true,
    baseURL: 'http://localhost:8000/'
});

export const userAPI = {
    getMe() {
        return instance.get("auth/discord/me")
            .then(response => {
                return response.data;
            })
            .catch(err => {
                return null;
            });
    },
    getCredits() {
        return instance.get("/credits")
            .then(response => {
                return response.data;
            })
            .catch(err => {
                return null;
            });
    },
    applyCredit(credit) {
        return instance.post("/credits/apply", credit)
            .then(response => {
                return response.data;
            })
            .catch(err => {
                return null;
            });
    }
}