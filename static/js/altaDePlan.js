$(document).ready(function() {
    getFormulariosYpisos()
});




async function getFormulariosYpisos() {
    const request = await fetch("http://3.83.152.157:5000/formularios", {
        method: 'GET',
    })

    const formularios = await request.json();
    console.log(formularios)
    if (request.ok) {
        let listadoHtml = '';
            for (let formulario of formularios) {
                listadoHtml +=
                `
                 <option value="${formulario.id}">${formulario.nombre}--${formulario.informacion}</option>
                `
        };

        document.getElementById("formulario").innerHTML = listadoHtml;

    }


    let obraid = null;
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    obraid = searchParams.get('obra_id');
    obra_id = Number(obraid);

    const request2 = await fetch("http://34.192.187.56:5000/pisos/" + obra_id, {
        method: 'GET',
    })

    const pisos = await request2.json();
    console.log(pisos)
    if (request.ok) {
        let listadoHtml = '';
            for (let piso of pisos) {
                listadoHtml +=
                `
                 <option value="${piso.id}">${piso.numero}</option>
                `
        };

        document.getElementById("piso").innerHTML = listadoHtml;

    }
}

async function registrarDocumento() {
    let datos = {};
    let obraid = null;
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    obraid = searchParams.get('obra_id');
    

    let Formulario = document.getElementById('formulario').value;
    datos.formulario_id = Number(Formulario);
    let Piso = document.getElementById('piso').value;
    datos.piso_id = Number(Piso);
    datos.obra_id = Number(obraid);
           const request = await fetch("http://34.192.187.56:5000/document/addDocument", {
               method: 'POST',
               body: JSON.stringify(datos),
               headers: {
                   'Accept': 'application/json',
                   'Content-Type': 'application/json'
               },
           })
         if (request.status == 201) {
             alert("documento registrado correctamente");
         }else{
             alert("Error registrando el documento");
         }
             window.location.href = `documentosDeObra.html?obra_id=${obraid}`;
   }