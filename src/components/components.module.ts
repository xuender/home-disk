import { NgModule } from '@angular/core';
import { IonicModule } from 'ionic-angular'
import { FileUploadModule } from 'ng2-file-upload'
//import { NgxUploaderModule } from 'ngx-uploader'

import { HdUploadComponent } from './hd-upload/hd-upload';
import { HdFilesComponent } from './hd-files/hd-files';
import { HdThumbnailComponent } from './hd-thumbnail/hd-thumbnail';
@NgModule({
  declarations: [
    HdUploadComponent,
    HdFilesComponent,
    HdThumbnailComponent,
  ],
  imports: [
    IonicModule,
    FileUploadModule,
  ],
  exports: [
    HdUploadComponent,
    FileUploadModule,
    HdFilesComponent,
    HdThumbnailComponent,
  ]
})
export class ComponentsModule { }
