$(document).ready(function() {
    getFormulariosYpisos()
});




async function getFormulariosYpisos() {
    const request = await fetch("http://localhost:5000/formularios", {
        method: 'GET',
    })

    const formularios = await request.json();
    console.log(formularios)
    if (request.ok) {
        let listadoHtml = '';
            for (let formulario of formularios) {
                listadoHtml +=
                `
                 <option value="${formulario.id_formulario}">${formulario.nombre}</option>
                `
        };

        document.getElementById("select").innerHTML = listadoHtml;

    }


    const request2 = await fetch("http://localhost:5000/pisos", {
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

        document.getElementById("select2").innerHTML = listadoHtml;

    }
}

async function registrarDocumento() {
    let datos = {};
    let obraid = null;
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    obraid = searchParams.get('obra_id');
    
    datos.formulario = document.getElementById('formulario').value;
    datos.piso = document.getElementById('piso').value;
    datos.obra = obraid;
           const request = await fetch("http://localhost:5000/formularios/registrar", {
               method: 'POST',
               body: JSON.stringify(datos),
               headers: {
                   'Accept': 'application/json',
                   'Content-Type': 'application/json'
               },
           })
         if (request.status == 200) {
             alert("documento registrado correctamente");
         }else{
             alert("Error registrando el documento");
         }
             window.location.href = `documentosDeObra.html`;
   }