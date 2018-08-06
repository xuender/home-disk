/**
 * 资源
 */
export interface File {
    /**
     * 主键
     */
    id: string
    /**
     * 类型
     */
    type: string
    /**
     * 名称
     */
    name: string
    /**
     * 创建时间
     */
    ct: Date
    /**
     * 子类型
     */
    sub: string
}
