$(document).ready(function() {
    getFormularios()
});




async function getFormularios() {
    const request = await fetch("http://34.192.187.56:5000/formularios", {
        method: 'GET',
    })

    const formularios = await request.json();
    console.log(formularios)
    if (request.ok) {
        let listadoHtml = '';
            for (let formulario of formularios) {
                listadoHtml +=
                `
                 <option value="${formulario.ID}">${formulario.Nombre}</option>
                `
        };

        document.getElementById("select").innerHTML = listadoHtml;

    }
}