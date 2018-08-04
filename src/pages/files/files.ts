import { Component } from '@angular/core';
import { NavController, NavParams } from 'ionic-angular';

@Component({
  selector: 'page-files',
  templateUrl: 'files.html',
})
export class FilesPage {
  days = [
    '2018-08-04',
    '2018-08-03',
    '2018-08-02',
    '2018-08-01',
    '2018-07-31',
    '2017-07-30',
    '2017-07-29',
    '2017-07-28',
    '2017-07-21',
    '2017-07-11',
    '2017-07-10',
  ]
  items = []
  constructor(public navCtrl: NavController, public navParams: NavParams) {
    this.items.push(...this.days.slice(0, 5))
  }
  doInfinite(infiniteScroll) {
    setTimeout(() => {
      if (this.items.length < this.days.length) {
        this.items.push(this.days[this.items.length])
      }
      infiniteScroll.complete();
    }, 500)
  }
}
