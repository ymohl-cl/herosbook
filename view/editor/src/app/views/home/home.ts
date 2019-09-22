import Vue from 'vue';
import Component from 'vue-class-component';
import HelloWorld from '../../components/HelloWorld/HelloWorld.vue';

import userService from '../../../services/user.service';
import navService from '../../../services/nav.service';

@Component({ 
    components: { HelloWorld },
})
export default class Home extends Vue {

    public mounted(){
        if(!userService.isConnected()){
            navService.replaceView('login');
        }
    }

}
