import Vue from 'vue';
import Component from 'vue-class-component';

import navService from '../services/nav.service';

@Component({
  name: 'App',
})
export default class App extends Vue {

  isDisplayAppBar():boolean{
    return navService.isSameRoute('home');
  }
}
