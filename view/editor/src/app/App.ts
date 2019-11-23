import Vue from 'vue';
import Component from 'vue-class-component';

import navService from '@/services/nav.service';
import userService from '@/services/user.service';

@Component({
  name: 'App',
})
export default class App extends Vue {
  public isDisplayAppBar():boolean {
    return navService.isSameRoute('home') || navService.isSameRoute('book');
  }

  public getPseudo():string {
    return userService.getPseudo();
  }
}
