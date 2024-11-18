import pymysql  #连接数据库
from flask import Flask,request,jsonify
from flask_cors import CORS

dp=pymysql.connect(host="localhost", user="root", password="123456", database="treehole")
cursor=dp.cursor()

app=Flask(__name__)
CORS(app,resources=r'/*')

@app.route('/list',methods=['POST'])
def list():
    if request.method=='POST':
        cursor.execute("SELECT * FROM user")
        data=cursor.fetchall()
        temp={}
        result=[]
        if(data!=None):
            for i in data:
                temp['username']=i[0]
                temp['password']=i[1]
                temp['userid']=i[2]
                result.append(temp.copy())
            print(result)
            return jsonify({'code':0,'data':result})
        else:
            return jsonify({'code':1,'message':'no data'})

@app.route('/add',methods=['POST'])
def add():
    if request.method=='POST':
        username=request.form.get('username')
        password=request.form.get('password')
        userid=request.form.get('userid')
        try:
            cursor.execute("INSERT INTO user(username,password,userid) VALUES(%s,%s,%s)",(username,password,userid))
            dp.commit()
            return jsonify({'message':'success'})
        except:
            return jsonify({'message':'failed'})
        
@app.route('/delete',methods=['POST'])
def delete():
    if request.method=='POST':
        userid=request.form.get('userid')
        try:
            cursor.execute("DELETE FROM user WHERE userid=%s",(userid))
            dp.commit()
            return jsonify({'message':'success'})
        except:
            return jsonify({'message':'failed'})
    
@app.route('/update',methods=['POST'])
def update():
    if request.method=='POST':
        userid=request.form.get('username')
        password=request.form.get('password')
        userid=request.form.get('userid')
        try:
            cursor.execute("UPDATE user SET password=%s,username=%s WHERE userid=%s",(password,userid))
            dp.commit()
            return jsonify({'message':'success'})
        except:
            return jsonify({'message':'failed'})
        
if __name__=='__main__':
    app.run(host='0.0.0.0',port=10500, debug=True)
