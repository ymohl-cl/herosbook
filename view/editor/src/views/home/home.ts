import Vue from 'vue';
import Component from 'vue-class-component';
import confirmDialog from '@/components/confirmDialog/confirmDialog.vue';


import userService from '@/services/user.service';
import navService from '@/services/ServiceNavigation';
import httpService from '@/services/http.service';
import session, {Session} from '@/services/ServiceSession';
import Book from '@/services/ControllerBook';

@Component({
  components: { confirmDialog },
})
export default class Home extends Vue {
  public session:Session = session;

  public getBooksLaunch:boolean = false;

  public validForm:boolean = false;

  public createNewBookDisplay:boolean = false;

  public createBookLaunch:boolean = false;

  public genres:any[] = [
    { key: 'fantasy', name: 'Fantaisie' },
    { key: 'scifi', name: 'Science Fiction' },
    { key: 'reallife', name: 'Dans la vie réelle' },
    { key: 'other', name: 'Autre' },
  ];

  public formCreateBook:any = {
    title: '',
    description: '',
    genre: '',
  };

  public titleRules:any = [
    (v:string) => !!v || 'Title is required',
  ];

  public genreRules:any = [
    (v:string) => !!v || 'Genre is required',
  ];

  public wantDeleteBookId:string = '';

  public deleteBookLaunch:boolean = false;

  public mounted() {
    if (!this.session.user.isConnected()) {
      navService.replaceView('/login');
    } else {
      this.getBooks();
    }
  }

  private getBooks() {
    this.getBooksLaunch = true;
    const headers = httpService.appendHeaders(httpService.getDefaultHeaders(), 'Authorization', `Bearer ${userService.getToken()}`);
    httpService.post('api/books/_searches', {}, headers, (response:any) => {
      this.session.books = response.data;
      this.getBooksLaunch = false;
    }, (error:any) => {
      this.getBooksLaunch = false;
    });
  }

  public goBook(idBook:string) {
    navService.changeView(`book/${idBook}`);
  }

  public isValid() {
    return this.validForm && !this.createBookLaunch;
  }

  public createBook() {
    this.createBookLaunch = true;
    let b = new Book("one title example", "one description example", "one genre exemple");
    this.session.addBook(b)
    this.createBookLaunch = false;
  }

  public openDeleteDialog(idBook:string) {
    this.wantDeleteBookId = idBook;
  }

  public deleteBookAction() {
    this.deleteBookLaunch = true;
    if (this.wantDeleteBookId !== '') {
      console.log(this.wantDeleteBookId);
      const headers = httpService.appendHeaders(httpService.getDefaultHeaders(), 'Authorization', `Bearer ${userService.getToken()}`);
      httpService.delete(`api/books/${this.wantDeleteBookId}`, headers, (response:any) => {
        this.deleteBookLaunch = false;
        this.session.removeBook(this.wantDeleteBookId)
        this.wantDeleteBookId = '';
      }, (error:any) => {
        this.deleteBookLaunch = false;
      });
    }
  }
}
