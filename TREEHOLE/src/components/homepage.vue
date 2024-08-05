<template>
    <div class="homepage">
        <div class="header"></div>
        <div class="headline">
            <h1>welcome to treehole！！!</h1>
            <p>You can say anything here,and nobody know who you are.</p>
        </div>
        
        <div class="container">
            <div class="set">
                <div class="personal">
                    <div class="avatar">
                        <img src="../image/original.png" alt="avatar">
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

</div>
    
</template>

<script>
    
</script>

<style scoped lang="less">
@import 'homepage.less';

</style>