import {defineStore} from 'pinia'
import axios from "axios";
import type {Login} from "~/types";

export const useUserStore = defineStore('token', () => {
    const user = ref();
    const token = useCookie("MY_COOKIE", {
        maxAge: 30,
    });
    const setToken = (data?: string) => (token.value = data);
    const setUser = (data?: any) => (user.value = data);

    const login = async (data: Login) => {
        try {
            const response = await axios.post('http://localhost:8080/user/auth', data);
            setToken(response.data.token)
        } catch (error) {
            setToken()
            setUser()
            console.log(error)
        }
    };


    const getUserData = async () => {
        if (token.value) {
            try {
                const response = await axios.get('http://localhost:8080/protected', {
                    headers: {
                        Authorization: `Bearer ${token.value}`
                    }
                })
                setUser(response.data)
            } catch (error) {
                setUser()
                console.log(error)
            }
        }
    }

    return {user, token, login, getUserData}
})