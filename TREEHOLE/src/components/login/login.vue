<template>
	<div class="entirety" v-if="visible">
	  <div class="background" @click="close"></div>
	  <div class="contain">
		<div class="big-box" :class="{ active: isLogin }">
		  <div class="big-contain" key="bigContainLogin" v-if="isLogin">
			<div class="btitle">LOGIN ACCOUNT</div>
			<div class="bform">
			  <input type="email" placeholder="Email" v-model="form.useremail" />
			  <span class="errTips" v-if="emailError"
				>*Email filled in incorrectly *</span
			  >
			  <input
				type="password"
				placeholder="Password"
				v-model="form.userpwd"
			  />
			  <span class="errTips" v-if="passwordError"
				>* Password filled in incorrectly *</span
			  >
			</div>
			<button class="bbutton" @click="login">LOGIN</button>
		  </div>
		  <div class="big-contain" key="bigContainRegister" v-else>
			<div class="btitle">CREATE ACCOUNT</div>
			<div class="bform">
			  <input type="text" placeholder="Username" v-model="form.username" />
			  <span class="errTips" v-if="existed"
				>*Username already exists!*</span
			  >
			  <input type="email" placeholder="Email" v-model="form.useremail" />
			  <input
				type="password"
				placeholder="Password"
				v-model="form.userpwd"
			  />
			</div>
			<button class="bbutton" @click="register">REGISTER</button>
		  </div>
		</div>
		<div class="small-box" :class="{ active: isLogin }">
		  <div class="small-contain" key="smallContainRegister" v-if="isLogin">
			<div class="stitle">Hi!</div>
			<p class="scontent">Are you ready to start your treehole journey?</p>
			<button class="sbutton" @click="changeType">REGISTER</button>
		  </div>
		  <div class="small-contain" key="smallContainLogin" v-else>
			<div class="stitle">Welcome back!</div>
			<p class="scontent">Please log in to your account</p>
			<button class="sbutton" @click="changeType">LOGIN</button>
		  </div>
		</div>
	  </div>
	</div>
  </template>
  
  <script>
  export default {
	name: "login",
	data() {
	  return {
		visible: true,
		isLogin: true,
		emailError: false,
		passwordError: false,
		existed: false,
		form: {
		  username: "",
		  useremail: "",
		  userpwd: "",
		},
	  };
	},
	methods: {
	  close() {
		this.$emit("close");
	  },
	  changeType() {
		this.isLogin = !this.isLogin;
		this.form.username = "";
		this.form.useremail = "";
		this.form.userpwd = "";
	  },
	  login() {
		if (this.form.useremail && this.form.userpwd) {
		  this.$axios
			.post("http://127.0.0.1:10520/api/user/login", {
			  email: this.form.useremail,
			  password: this.form.userpwd,
			})
			.then((res) => {
			  if (res.data === 0) {
				alert("登录成功!");
				this.close();
			  } else if (res.data === -1) {
				this.emailError = true;
			  } else if (res.data === 1) {
				this.passwordError = true;
			  }
			})
			.catch((err) => {
			  console.error(err);
			});
		} else {
		  alert("字段不能为空！");
		}
	  },
	  register() {
		if (this.form.username && this.form.useremail && this.form.userpwd) {
		  this.$axios
			.post("http://127.0.0.1:10520/api/user/add", {
			  username: this.form.username,
			  email: this.form.useremail,
			  password: this.form.userpwd,
			})
			.then((res) => {
			  if (res.data === 0) {
				alert("Registration successful!");
				this.changeType();
			  } else if (res.data === -1) {
				this.existed = true;
			  }
			})
			.catch((err) => {
			  console.error(err);
			});
		} else {
		  alert("填写不能为空!");
		}
	  },
	},
  };
  </script>
  
  <style scoped lang="less">
  @import "login.less";
  </style>