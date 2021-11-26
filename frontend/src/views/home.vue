<template>
    <div>
        <Header/>
        <h1>Trabajo Final de Programacion Concurrente y Distribuida</h1>
        <h4>Este registro sirve para los casos de los agresores sentenciados por hechos de violencia en el marco de la Ley Nº 30364</h4>
        <br>
        <router-link to="/">Home</router-link> 
        <router-link to="/dashboard" style="margin-left: 10px">Tabla</router-link>
        <br>
        <h4>Docente de la materia:</h4>
        <br>
        <h6>Carlos Alberto Jara Garcia</h6>
        <br>
        <h4>Seccion</h4>
        <br>
        <h6>- SW72</h6>
        <br>
        <h4>Integrantes de equipo:</h4>
        <br>
        <h6>- Juben Antony Cacha Panduro</h6>
        <br>
        <h6>- Gerardo Josue Huerta Macedo</h6>
        <br>
        <h6> - Daurih Paulo Alegre Rocha</h6>
        <br>
        <h6>- Jesus Fernando Olano Pereda</h6>
        <br>
        <h6>- Denilson Ramirez Zarate Denilson</h6>
        <br>
        <Footer/>
    </div>
</template>

<script>

import Header from '@/components/Header.vue';
import Footer from '@/components/Footer.vue';
//import axios from 'axios';
export default{
    name:"Home",
    setup() {
        
    },
    components:{
        Header,
        Footer
    },
}
</script>


<style  scoped>
    .izquierda{
        text-align: left;
    }
</style>