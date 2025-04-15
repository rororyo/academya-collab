"use strict";

const input_pdf = document.getElementById('cv');
const upload_cv_btn = document.getElementById('upload-cv');
const hasil = document.getElementById('hasil');

hasil.style.display = 'None';

upload_cv_btn.addEventListener('click',async function(event) {
    const formData = new FormData();
    const file = input_pdf.files[0];
    formData.append('pdf_file',file);
    try{
        const response = await fetch('http://127.0.0.1:5000/predict',{
            method: 'POST',
            body: formData
        });
        const result = await response.json();
        console.log(result);
        hasil.style.display = 'block';
        hasil.textContent = `Kamu cocoknya jadi ${result.prediksi}`;
    }catch(err){
        console.log(err)
    }
})