import Vue from 'vue';
import Component from 'vue-class-component';
import HttpService from '../../../services/http.service';

/**
 * First component : reutilisable node
 */
@Component
export default class HelloWorld extends Vue {

    constructor(){
        super();
    }

    ecosystem = [{
            text: 'vuetify-loader',
            href: 'https://github.com/vuetifyjs/vuetify-loader',
        }, {
            text: 'github',
            href: 'https://github.com/vuetifyjs/vuetify',
        }, {
            text: 'awesome-vuetify',
            href: 'https://github.com/vuetifyjs/awesome-vuetify',
        },
    ];
    importantLinks = [{
            text: 'Documentation',
            href: 'https://vuetifyjs.com',
        }, {
            text: 'Chat',
            href: 'https://community.vuetifyjs.com',
        }, {
            text: 'Made with Vuetify',
            href: 'https://madewithvuejs.com/vuetify',
        }, {
            text: 'Twitter',
            href: 'https://twitter.com/vuetifyjs',
        }, {
            text: 'Articles',
            href: 'https://medium.com/vuetify',
        },
    ];
    whatsNext = [{
            text: 'Explore components',
            href: 'https://vuetifyjs.com/components/api-explorer',
        }, {
            text: 'Select a layout',
            href: 'https://vuetifyjs.com/layout/pre-defined',
        }, {
            text: 'Frequently Asked Questions',
            href: 'https://vuetifyjs.com/getting-started/frequently-asked-questions',
        },
    ];


    testAxios(){
        HttpService.get('/ping', (response:any) => {
            console.log(response.data)
        }, (error:any) => {
            console.error(error);
        });

    }
}