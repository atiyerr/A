<template>
	<button class="close"></button>
	<div class="login-register" >	
		<div class="contain">
			<div class="big-box" :class="{active:isLogin}">
				<div class="big-contain" key="bigContainLogin" v-if="isLogin">
					<div class="btitle">LOGIN ACCOUNT</div>
					<div class="bform">
						<input type="email" placeholder="Email" v-model="form.useremail">
						<span class="errTips" v-if="emailError">*Email filled in incorrectly *</span>
						<input type="password" placeholder="Password" v-model="form.userpwd">
						<span class="errTips" v-if="emailError">* Password filled in incorrectly *</span>
					</div>
					<button class="bbutton" @click="login">LOGIN</button>
				</div>
				<div class="big-contain" key="bigContainRegister" v-else>
					<div class="btitle">CREATE ACCOUNT</div>
					<div class="bform">
						<input type="text" placeholder="Username" v-model="form.username">
						<span class="errTips" v-if="existed">*Username already exists!*</span>
						<input type="email" placeholder="Email" v-model="form.useremail">
						<input type="password" placeholder="Password" v-model="form.userpwd">
					</div>
					<button class="bbutton" @click="register">REGISTER</button>
				</div>
			</div>
			<div class="small-box" :class="{active:isLogin}">
				<div class="small-contain" key="smallContainRegister" v-if="isLogin">
					<div class="stitle">Hi!</div>
					<p class="scontent">Are you ready to start your treehole journey?</p>
					<button class="sbutton" @click="changeType">REGISTER</button>
				</div>
				<div class="small-contain" key="smallContainLogin" v-else>
					<div class="stitle">Welcome back!</div>
					<p class="scontent">Please log in to your account</p>
					<button class="sbutton" @click="changeType">login</button>
				</div>
			</div>
		</div>
	</div>
</template>

<script>

	// 监听窗口滚动事件，动态调整login组件的位置
	window.addEventListener('scroll', () => {
    	const loginElement = document.querySelector('.login-register');
    	if (loginElement) {
      	  const scrollTop = window.scrollY;
      		loginElement.style.top = `${scrollTop}px`;
    	}
  	});

	export default {
		name:'login-register',
		data(){
			return {
				isLogin:false,
				emailError: false,
				passwordError: false,
				existed: false,
				form:{
					username:'',
					useremail:'',
					userpwd:''
				}
			}
		},
		methods:{
			changeType() {
				this.isLogin = !this.isLogin
				this.form.username = ''
				this.form.useremail = ''
				this.form.userpwd = ''
			},
			login() {
				const self = this;
				if (self.form.useremail != "" && self.form.userpwd != "") {
					self.$axios({
						method:'post',
						url: 'http://127.0.0.1:10520/api/user/login',
						data: {
							email: self.form.useremail,
							password: self.form.userpwd
						}
					})
					.then( res => {
						switch(res.data){
							case 0: 
								alert("登陆成功！");
								break;
							case -1:
								this.emailError = true;
								break;
							case 1:
								this.passwordError = true;
								break;
						}
					})
					.catch( err => {
						console.log(err);
					})
				} else{
					alert("填写不能为空！");
				}
			},
			register(){
				const self = this;
				if(self.form.username != "" && self.form.useremail != "" && self.form.userpwd != ""){
					self.$axios({
						method:'post',
						url: 'http://127.0.0.1:10520/api/user/add',
						data: {
							username: self.form.username,
							email: self.form.useremail,
							password: self.form.userpwd
						}
					})
					.then( res => {
						switch(res.data){
							case 0:
								alert("注册成功！");
								this.login();
								break;
							case -1:
								this.existed = true;
								break;
						}
					})
					.catch( err => {
						console.log(err);
					})
				} else {
					alert("填写不能为空！");
				}
			}
		}
	}
</script>


<style scoped lang="less">

    @import 'login.less';
    
</style>