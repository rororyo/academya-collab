import re
from transformers import pipeline,AutoModelForSequenceClassification, AutoTokenizer
import pandas as pd
import matplotlib.pyplot as plt
import torch
from flask import Flask, request, jsonify
from flask_cors import CORS

import google.generativeai as genai
from google.generativeai import types, configure
import pymupdf

app = Flask(__name__)
CORS(app,origins=["http://localhost:4000"])


API_KEY = 'AIzaSyCT84o5ZsloW0Aa5dNlq_k4IIE-XZFlGNU'
pendahuluan = 'Kamu adalah model AI tingkat tinggi yang dirancang untuk menganalisis pekerjaan yang cocok untuk seseorang berdasarkan CV nya. Kamu akan menerima CV yang sudah diubah menjadi string. Tugasmu adalah memberikan nama pekerjaan yang cocok untuk orang tersebut.\nOutput yang kamu berikan harus berupa nama pekerjaan yang paling cocok dengan orang tersebut. Berikut adalah cv nya\n'

genai.configure(api_key=API_KEY)
def prediksi(text):
    model = genai.GenerativeModel("gemini-1.5-flash")
    response = model.generate_content(pendahuluan+text)
    return response

def pipeline(cv):
    pdf = pymupdf.open(stream=cv,filetype='pdf')
    text = chr(12).join([page.get_text() for page in pdf])
    print(text)
    response = prediksi(text)
    
    return response.text

@app.route('/predict',methods=['POST'])
def predict():
    if "pdf_file" not in request.files:
        return jsonify({'error':'Tidak ada file'}),400
    file = request.files['pdf_file']
    if file.filename == '':
        return jsonify({'error':'Tidak ada file'}),400
    cv = file.read()
    hasil_prediksi = pipeline(cv)
    return jsonify({'prediksi': hasil_prediksi})

if __name__ == '__main__':
    app.run(debug=True)