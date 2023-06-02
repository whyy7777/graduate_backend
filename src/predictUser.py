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
deleteSql = "DELETE FROM music_online.recommend_songs WHERE userId = " + userId + ";"
engine.execute(deleteSql)
selectSql = "SELECT id FROM music_online.songs WHERE category = " + str(int(userType)) + ";"
print(selectSql)
data = pd.read_sql_query(selectSql, engine)
for i in range(0, len(data)):
    insertSql = "INSERT INTO music_online.recommend_songs (songId, userId) VALUES (" + str(int(data.loc[i])) + ", " + userId + ");"
    print(insertSql)
    engine.execute(insertSql)
