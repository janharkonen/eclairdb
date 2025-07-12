<template>
  <div class="w-full h-full flex flex-col">
    <div class="logo-div">
      <img 
      src="https://picapi.janharkonen.fi/api/pics/ba7cbf1c9780435899c83df67d5a1747.png"
      alt="EclairDB Logo"
      class="logo-img" 
      />
    </div>
    <div class="inputs-div">
      <div class="input-div">
        <img src="https://upload.wikimedia.org/wikipedia/commons/2/29/Postgresql_elephant.svg"
        alt="Postgres Logo" 
        class="w-8 h-8 flex-shrink-0" 
        />
        <input 
        type="text" 
        placeholder="postgres://user_name:pass_word@host-IP:port/database-name?sslmode=disable" 
        class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
        />
      </div>
    </div>
  </div>
</template>

<style scoped>
.logo-div {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 2rem;
}
.logo-img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  max-width: 240px;
  max-height: 240px;
}
.inputs-div {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  padding-top: 8rem;
  gap: 2rem;
  width: 100%;
  height: 100%;
  flex-grow: 1;
}
.input-div {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  gap: 1rem;
  width: 100%;
  max-width: 920px;
}
</style>