import joblib
import pandas as pd
from sklearn.model_selection import train_test_split
gb2 = joblib.load('saved_model/gbSong.pkl')
df = pd.read_csv('/Users/yiny/Desktop/graduate/songs.csv')
X = df.drop('category', axis=1)
y = df['category']
X = pd.get_dummies(X, drop_first=True)
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.30)
print(gb2.predict(X_test))
