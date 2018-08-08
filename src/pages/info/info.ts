import { Component } from '@angular/core';
import { NavParams, ViewController } from 'ionic-angular';

@Component({
  selector: 'page-info',
  templateUrl: 'info.html',
})
export class InfoPage {

  file: File
  constructor(
    public navParams: NavParams,
    public viewCtrl: ViewController,
  ) {
    this.file = navParams.get('file')
  }
  cancel() {
    this.viewCtrl.dismiss()
  }
  ionViewDidLoad() {
    console.log('ionViewDidLoad InfoPage');
  }

}
