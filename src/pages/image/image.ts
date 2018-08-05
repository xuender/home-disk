import { Component } from '@angular/core';
import {  NavParams,ViewController } from 'ionic-angular';
import {File} from '../../domain/file'
@Component({
  selector: 'page-image',
  templateUrl: 'image.html',
})
export class ImagePage {
  file: File
  constructor(
    public navParams: NavParams,
    public viewCtrl: ViewController,
  ) {
    this.file = navParams.get('file')
  }
  cancel(){
    this.viewCtrl.dismiss()
  }
  ionViewDidLoad() {
    console.log('ionViewDidLoad ImagePage');
  }
}
