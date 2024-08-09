<template>
    <div class="below">
        <div class="above">
            <div class="logo">
                <img src="../../image/logo.png" alt="logo">
                <h3>TREEHOLE</h3>
            </div>
            <ul class="variety">
                <li><a href="#friends">friends</a></li>
                <li><a href="#set">set</a></li>
            </ul>
            <ul class="state">
                <li ><a href="#login" @click.prevent="Login">login</a></li>
                <li><a href="#logout">logout</a></li>
                <li><a href="#register">register</a></li>
            </ul>
        </div>
        <div class="homepage">
            <h1>welcome to treehole！！</h1>
            <p>You can say anything here,and nobody know who you are.</p>
        
            <div class="container">
            <div class="set">
                <div class="personal">
                    <div class="avatar">
                        <img src="../../image/original.png" alt="avatar">
                    </div>
                    <ul class="profile">
                        <li>username：{{username}}</li>
                        <li>current identity：<a href="#change">identity-></a></li>
                        <li>change profile:<a href="#edit">edit</a></li>
                    </ul>
                </div>
                <div class="chat">聊天区</div>
            </div>
            
            <div class="comment">
                <div class="publish">
                    select theme：<select>
                        <option value="career">career</option>
                        <option value="sentiment">sentiment</option>
                        <option value="complain">complain</option>
                        <option value="others">others</option>
                    </select>
                    <textarea class="text" placeholder="say something..."></textarea>
                    <button type="submit" class="btn">publish</button>
                    <button class="btn">emoji</button>
                    <label class="file-label">
                        <span>Upload Files</span>
                        <input type="file" multiple style="display: none;">
                    </label>
                </div>
                <div class="comment-container">
                    <ul class="comment-title">
                        <li><a href="#all">all</a></li>
                        <li><a href="#career">career</a></li>
                        <li><a href="#sentiment">sentiment</a></li>
                        <li><a href="#complain">complain</a></li>                    
                        <li><a href="#others">others</a></li>
                    </ul>
                </div>
                <div class="comment-list">评论区</div>
            </div>
            </div>

            <div class="footer">
            <ul class="links">
                <li><a href="#about">About us</a></li>
                <li>|</li>
                <li><a href="#contact">Contact us</a></li>
            </ul>
            </div>
        </div> 
    </div>
    
</template>

<script></script>
   

<style scoped lang="less">


    // 适配不同大小窗口
    @import 'homepage1.less';
    @media (max-width: 939px) {
        @import 'homepage2.less';
    }
</style>