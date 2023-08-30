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
}