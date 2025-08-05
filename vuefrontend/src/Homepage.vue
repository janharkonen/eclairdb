<template>
  <div class="
    w-full
    opacity-90
    max-w-5xl
    mx-auto
    shadow-2xl
    border-2
    border-cyan-300
    flex 
    flex-col 
    bg-gradient-to-tl from-cyan-800 to-cyan-700
    rounded-xl 
    px-4 
    pb-10
    font-mono 
    text-gray-100 
    text-xs 
    md:text-base">
    <div class="logo-div"> 
      <img 
      src="https://picapi.janharkonen.fi/api/pics/7b99c61a71c741c3a51896123e8069db.png"
      alt="EclairDB Logo"
      class="logo-img opacity-100"
      style="filter: invert(1);"
      />
    </div>
    <div class="inputs-div">
      <DatabaseInput
        img-src="https://upload.wikimedia.org/wikipedia/commons/2/29/Postgresql_elephant.svg"
        service="PostgreSQL"
        placeholder="postgres://user_name:pass_word@host-IP:port/database-name?sslmode=require&channel_binding=require"
        apiroute="connect-postgres"
      />
      <!--
        <DatabaseInput
        img-src="https://picapi.janharkonen.fi/api/pics/c6bed4194e4341fc992e13a0487fb329.png"
        service="SQLite"
        placeholder="TO BE IMPLEMENTED"
        apiroute="connect-sqlite"
        />
      -->
    </div>
  </div>
</template>

<script setup lang="ts">
import DatabaseInput from './components/DatabaseInput.vue'



</script>

<style scoped>
.logo-div {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  gap: 2rem;
}
.logo-img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  max-width: 320px;
  max-height: 320px;
}
.inputs-div {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  width: 100%;
}
</style>