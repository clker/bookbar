<template>
  <div id="">
    <div class="layout">
        <Menu mode="horizontal" theme="dark" active-name="1">
            <div class="layout-logo"><a href="/#">BookBar</a></div>
            <div class="layout-nav">
                <MenuItem name="1"><a href="/#">Books</a> </MenuItem>
                <MenuItem name="2"> Topics </MenuItem>
                <MenuItem name="3"> Signin </MenuItem>
            </div>
        </Menu>
    </div>

    <router-view></router-view>
  </div>
</template>



<script>
export default {
  name: 'app',
  methods : {
  }
}
</script>

<style>
    .layout-logo {
        margin-left:50px;
        width: 300px;
        height: 30px;
        //background: #5b6270;
        float: left;
        //position: relative;
        //top: 15px;
        //left: 20px;
        color: #FFFFFF;
        font-size: 28px;
    }

    .layout-nav {
        float:right;
        margin-right:50px;
    }

    .layout-ceiling {
        background: #f5f7f9;
        padding: 5px 50px;
    }

    .layout-ceiling-main {
        float: right;
    }

    .layout-ceiling-main * {
        margin-right:20px;
    }

    .ivu-menu-dark a {
        color: #FFFFFF;
    }

    .mt-bg {
        margin-top: 50px;
    }

    .mt-md {
        margin-top: 20px;
    }

    .aright {
    }

    .clearfix {
        clear:both;
    }

    .container {
        margin: 0 50px 0 50px;
    }

    .h1 {
        font-size: 28px; font-weight: bold; color: rgb(70, 76, 91);
    }

    .list-style-fix ul li {
        list-style-type: disc
    }

    .list-style-fix ol li {
        list-style-type: decimal
    }


    body {
        font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
        font-size: 14px; color: rgb(70, 76, 91);
    }

</style>
