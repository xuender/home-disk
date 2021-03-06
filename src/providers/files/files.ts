import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { File } from '../../domain/file'
import { Observable } from 'rxjs/Observable';
import { concatAll } from 'rxjs/operators'
import { of } from 'rxjs/observable/of'
import { filter, includes } from 'lodash-es'
@Injectable()
export class FilesProvider {
  private _days: string[] = []
  filesMap = new Map<string, Array<File>>()
  days = []
  private run = true
  news: File[] = []
  constructor(public http: HttpClient) {
    this.reset()
  }
  reset() {
    this.run = true
    this.days = []
    this.news = []
    this.filesMap.clear()
    this.http.get('/days')
      .subscribe((days: string[]) => {
        // 获取所有日期
        this._days = days
        // 初始化 5 天
        const td = this._days.splice(0, 5)
        if (td.length == 0) {
          this.run = false
          return
        }
        // 获取 5 天文件列表
        const s = [];
        for (const d of td) {
          s.push(this.getFiles(d))
        }
        // 同步执行
        const source = of(...s);
        source.pipe(concatAll())
          .subscribe(files => {
            const d = td.shift()
            console.log('reset d:', d)
            this.filesMap.set(d, files)
            this.days.push(d)
            if (td.length == 0) {
              this.run = false
            }
          })
      })
  }
  filter(d: string, types: string[]): File[] {
    if (types.length == 0) {
      return this.filesMap.get(d)
    }
    return filter(this.filesMap.get(d), f => includes(types, f.type))
  }
  getFiles(day: string): Observable<Array<File>> {
    return this.http.get<Array<File>>(`/days/${day}`)
  }
  loadDay(): Promise<boolean> {
    console.log('run', this.run)
    return new Promise<boolean>(resolve => {
      if (!this.run && this._days.length > 0) {
        this.run = true
        const d = this._days.shift()
        if (d) {
          console.log('load d:', d)
          this.getFiles(d).subscribe(files => {
            this.filesMap.set(d, files)
            this.days.push(d)
            this.run = false
            resolve(true)
          })
        } else {
          this.run = false
        }
      }
      resolve(true)
    })
  }
}
