<template>
  <div class="overlay hidden">
    <section class="popup hidden">
      <h2 class="popup__title">Введите имя файла</h2>
      <input class="popup__input" placeholder="Имя файла" required v-model="filename"/>
      <button class="popup__create" @click="getUserData">Создать</button>
      <button class="popup__close" @click="closePopup"><img class="popup__close_img" alt="Close"
                                                            src="../assets/images/Krest.svg"></button>
    </section>
  </div>
</template>


<script setup lang="ts">
import axios from "axios";

const userStore = useUserStore()
const router = useRouter()
const filename = ref('');

const getUserData = async () => {
  await userStore.getUserData()
  await router.push('/executor')
}

const createNewSource = async () => {
  const userId = userStore.user.id
  const token = userStore.token
  try {
    const response = await axios.post('http://localhost:8080/source/new', {
      headers: {
        Authorization: `Bearer ${token}`
      },
      body: {
        user_id: userId,
        filename: filename,
        data: ''
      }
    })
    console.log(response)
  } catch (error) {
    console.log(error)
  }
}

const closePopup = () => {
  const popupElement = document.querySelector('.popup');
  if (popupElement)
    popupElement.classList.add('hidden');

  const overlayElement = document.querySelector('.overlay');
  if(overlayElement)
    overlayElement.classList.add('hidden');

}
</script>

<style>
@import "../assets/createPopup.css";
</style>