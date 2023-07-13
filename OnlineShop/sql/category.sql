DROP TABLE IF EXISTS `tb_category`;
CREATE TABLE `tb_category` (
                               `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '类目id',
                               `name` varchar(20) NOT NULL COMMENT '类目名称',
                               `parent_id` bigint(20) NOT NULL COMMENT '父类目id,顶级类目填0',
                               `is_parent` tinyint(1) NOT NULL COMMENT '是否为父节点，0为否，1为是',
                               `sort` int(4) NOT NULL COMMENT '排序指数，越小越靠前',
                               PRIMARY KEY (`id`),
                               KEY `key_parent_id` (`parent_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1424 DEFAULT CHARSET=utf8 COMMENT='商品类目表，类目和商品(spu)是一对多关系，类目与品牌是多对多关系';
# 支持三级类目0,1,2

    # 0 级
INSERT INTO `tb_category` VALUES ('1', '图书、音像、电子书刊', '0', '1', '1');
    # 1 级
INSERT INTO `tb_category` VALUES ('2', '电子书刊', '1', '1', '1');
    # 2 级
# 2 的子类目
INSERT INTO `tb_category` VALUES ('3', '电子书', '2', '0', '1');
INSERT INTO `tb_category` VALUES ('4', '网络原创', '2', '0', '2');
INSERT INTO `tb_category` VALUES ('5', '数字杂志', '2', '0', '3');
    # 1 级
# 1 的子类目
INSERT INTO `tb_category` VALUES ('7', '音像', '1', '1', '2');
    # 1 级0

INSERT INTO `tb_category` VALUES ('11', '英文原版', '1', '1', '3');

    # 0 级
INSERT INTO `tb_category` VALUES ('103', '家用电器', '0', '1', '3');
    # 1 级
INSERT INTO `tb_category` VALUES ('104', '大 家 电', '103', '1', '1');
    # 2 级
INSERT INTO `tb_category` VALUES ('105', '平板电视', '104', '0', '1');
INSERT INTO `tb_category` VALUES ('106', '空调', '104', '0', '2');
INSERT INTO `tb_category` VALUES ('107', '冰箱', '104', '0', '3');
