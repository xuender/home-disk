import { Injectable, EventEmitter } from '@angular/core';
import { ModalController } from 'ionic-angular';
import { File } from '../../domain/file';
import { ImagePage } from '../../pages/image/image';
import { InfoPage } from '../../pages/info/info';

const PATH = 'assets/imgs/'
@Injectable()
export class PreviewProvider {

  srcDefault = `${PATH}unknown.png`
  itemChanged = new EventEmitter<any>();
  constructor(
    private modalCtrl: ModalController,
  ) {
    console.log('Hello PreviewProvider Provider');
  }
  preview(file: File) {
    this.modalCtrl
      .create(this.getPage(file), { file: file })
      .present();
  }
  private getPage(file: File): any {
    switch (file.type) {
      case 'image':
        return ImagePage
      default:
        return InfoPage
    }
  }
  getSrc(file: File): string {
    if (file) {
      switch (file.type) {
        case 'image':
          return `/t/${file.id}`
        case 'archive':
        case 'audio':
        case 'video':
          return `${PATH}${file.type}.png`
        default:
          return this.srcDefault
      }
    }
    return `${PATH}wait.png`
  }
}
