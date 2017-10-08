<template>
  <div id="wrapper">
    <nav class="navbar navbar-default">
      <div class="container">
        <a class="navbar-brand" href="#">
          <i class="glyphicon glyphicon-time"></i>
          BOOKBAR 
        </a>
        <ul class="nav navbar-nav">
          <li><router-link to="/books">Books</router-link></li>
          <li><router-link to="/topics">Topics</router-link></li>
        </ul>
        <ul class="nav navbar-nav navbar-right">
          <li><router-link to="/signup">Signup</router-link></li>
          <li><router-link to="/login">Login</router-link></li>
        </ul>
      </div>
    </nav>
    <div class="container">
      <!-- <div class="col-sm-3">
       <sidebar></sidebar> 
      </div>
      <div class="col-sm-9">
      </div>
      -->
      <router-view></router-view>
    </div>
  </div>
</template>



<script>
export default {
  name: 'app'
}
</script>

