$(document).ready(function() {
    getFormularios()
});




async function getFormularios() {
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