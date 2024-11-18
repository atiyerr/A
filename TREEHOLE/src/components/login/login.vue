<template>
	<div class="entirety" v-if="visible">
		<div class="background" @click="close"></div>
		<div class="contain">
			<div class="big-box" :class="{ active: isLogin }">
				<div class="big-contain" key="bigContainLogin" v-if="isLogin">
					<div class="btitle">LOGIN ACCOUNT</div>
					<div class="bform">
						<input type="userid" placeholder="Email/phone" v-model="form.userid" />
						<span class="errTips" v-if="idError">*userid filled in incorrectly *</span>
						<input type="password" placeholder="Password" v-model="form.password" />
						<span class="errTips" v-if="passwordError">* Password filled in incorrectly *</span>
					</div>
					<button class="bbutton" @click="login">LOGIN</button>
				</div>
				<div class="big-contain" key="bigContainRegister" v-else>
					<div class="btitle">CREATE ACCOUNT</div>
					<div class="bform">
						<!-- <input type="text" placeholder="Username" v-model="form.username" /> -->
						<!-- <span class="errTips" v-if="existed">*Username already exists!*</span> -->
						<input type="text" placeholder="Email/phone" v-model="form.userid" />
						<input type="password" placeholder="Password" v-model="form.password" />
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
import axios from 'axios';
export default {

	name: "login",
	data() {
		return {
			visible: true,
			isLogin: true,
			idError: false,
			passwordError: false,
			existed: false,
			form: {
				userid: "",
				password: "",
			},
		};
	},
	methods: {
		close() {
			this.$emit("close");
		},

		changeType() {
			this.isLogin = !this.isLogin;
			this.form.userid = "";
			this.form.password = "";
		},
		login() {
			if (this.form.userid && this.form.password) {
				axios
					.post("http://192.168.1.122:10500/user/load", {
						userid: this.form.userid,
						password: this.form.password,
					})
					.then(response => response.data)
					.then((data) => {
						console.log('login', data.code)

						if (data.code === '0') {
							alert("Login successful!");
							window.location.href = 'homepage.html';
							this.close();
						} else if (data.code === '2') {
							alert("User does not exist!");
							this.idError = true;
						} else if (data.code === '1') {
							this.passwordError = true;
							alert("Password error!");
						} else if(data.code === '3'){
							alert("User does not exist!");
							this.idError = true;
							this.passwordError = true;
                	}
				})
					.catch((error) => {
    					console.error("Error details:", error)
					});
			} else {
				alert("字段不能为空！");
			}
		},
		register() {
			if ( this.form.userid && this.form.password) {
				axios
					.post("http://192.168.1.122:10500/user/add", {
						userid: this.form.userid,
						password: this.form.password,
					})
					.then(response => response.data)
					.then((data) => {
						if (data.code === '0') {
							alert("Registration successful!");
							this.changeType();
						} else if (data.code === '1') {
							this.existed = true;
							alert("Registration failed!");
						}else{
							alert("unknown error");
						}
					})
					.catch((error) => {
						console.error(error);
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