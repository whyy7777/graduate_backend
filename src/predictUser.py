import joblib
import pandas as pd
from sklearn.model_selection import train_test_split
from sqlalchemy import create_engine
import pymysql
import sys

userId = sys.argv[1]
print(type(userId))
MYSQL_HOST = '1.117.65.130'
MYSQL_PORT = '3306'
MYSQL_USER = 'root'
MYSQL_PASSWORD = '613181Hyy.'
MYSQL_DB = 'music_online'
engine = create_engine('mysql+pymysql://%s:%s@%s:%s/%s?charset=utf8'
                       % (MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DB))

gb2 = joblib.load('saved_model/gbUser.pkl')

sql = "SELECT * FROM listen_record WHERE userId = " + userId + ";"
result = pd.read_sql_query(sql, engine)
result = result.drop(columns='userType')
userType = gb2.predict(result)
updateSql = "UPDATE music_online.listen_record t SET t.userType = " + str(int(userType)) + " WHERE t.userId = " + userId + ";"
# print(updateSql)
engine.execute(updateSql)
