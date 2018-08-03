import { Component } from '@angular/core';
import { NavController, NavParams } from 'ionic-angular';

@Component({
  selector: 'page-files',
  templateUrl: 'files.html',
})
export class FilesPage {
  constructor(public navCtrl: NavController, public navParams: NavParams) {
  }
}
