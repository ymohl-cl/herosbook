import Vue from 'vue';
import Component from 'vue-class-component';
import confirmDialog from '../../components/confirmDialog/confirmDialog.vue';


import userService from '@/services/user.service';
import navService from '@/services/nav.service';
import httpService from '@/services/http.service';

@Component({
  components: { confirmDialog },
})
export default class Home extends Vue {
  public books:any[] = [];

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
    if (!userService.isConnected()) {
      navService.replaceView('/login');
    } else {
      this.getBooks();
    }
  }

  private getBooks() {
    this.getBooksLaunch = true;
    const headers = httpService.appendHeaders(httpService.getDefaultHeaders(), 'Authorization', `Bearer ${userService.getToken()}`);
    httpService.post('api/books/_searches', {}, headers, (response:any) => {
      this.books = response.data;
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
    const headers = httpService.appendHeaders(httpService.getDefaultHeaders(), 'Authorization', `Bearer ${userService.getToken()}`);
    httpService.post('api/books', this.formCreateBook, headers, (response:any) => {
      this.createBookLaunch = false;
      this.createNewBookDisplay = false;
      this.goBook(response.data.identifier);
    }, (error:any) => {
      this.createBookLaunch = false;
    });
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
        this.books = this.books.filter(element => element.identifier !== this.wantDeleteBookId);
        this.wantDeleteBookId = '';
      }, (error:any) => {
        this.deleteBookLaunch = false;
      });
    }
  }
}
