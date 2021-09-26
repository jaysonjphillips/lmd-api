<template>
<div>
  <nav class="navbar navbar-expand-xl navbar-dark bg-dark" aria-label="Sixth navbar example">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">Inventory Management</a>
      <button class="navbar-toggler"
        type="button" data-bs-toggle="collapse"
        data-bs-target="#navbarsExample06"
        aria-controls="navbarsExample06"
        aria-expanded="false"
        aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse justify-content-end" id="navbarsExample06">
        <ul class="navbar-nav mb-2 mb-xl-0">
          <li class="nav-item">
            <router-link to="/" class="nav-link">Dashboard</router-link>
          </li>
          <li class="nav-item">
            <router-link to="/wine" class="nav-link">Wine</router-link>
          </li>
          <li class="nav-item">
            <router-link to="/" class="nav-link">Spirits</router-link>
          </li>
          <li class="nav-item">
            <router-link to="/" class="nav-link">Sneakers</router-link>
          </li>
          <li class="nav-item">
            <router-link to="/" class="nav-link">Moving Boxes</router-link>
          </li>
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" id="dropdown06"
               data-bs-toggle="dropdown" aria-expanded="false">Dropdown</a>
            <ul class="dropdown-menu" aria-labelledby="dropdown06">
              <li><router-link to="/" class="nav-link">Home</router-link></li>
              <li><router-link to="/" class="nav-link">Home</router-link></li>
              <li><router-link to="/" class="nav-link">Home</router-link></li>
            </ul>
          </li>
        </ul>
      </div>
    </div>
  </nav>
  <router-view/>
</div>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
