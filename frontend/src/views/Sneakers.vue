<template>
<div class="justify-content-start mt-4 mb-4">
    <h2 class="mx-4">Sneakers</h2>
     <WineList msg="Welcome to Your Vue.js + TypeScript App"/>
</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import WineList from '../components/WineList.vue'; // @ is an alias to /src

export default defineComponent({
  name: 'Sneakers',
  components: {
    WineList,
  },
});
</script>
