import pandas as pd
import joblib

from sklearn.model_selection import train_test_split
from sklearn.ensemble import GradientBoostingClassifier
from sklearn import svm
df = pd.read_csv('/Users/yiny/Desktop/graduate/songs.csv')
X = df.drop('category', axis=1)
y = df['category']
X = pd.get_dummies(X, drop_first=True)
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.30)
gb_model = GradientBoostingClassifier()
gb_model.fit(X_train, y_train)
gb_model.predict(X_test)
joblib.dump(gb_model, 'saved_model/gbSong.pkl')
