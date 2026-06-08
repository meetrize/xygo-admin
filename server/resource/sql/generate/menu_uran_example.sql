-- 悠然示例 菜单 SQL
-- 上级菜单ID: 0
-- ======= 顶级模式：创建目录(type=1) + 页面(type=2) + 按钮(type=3) =======

-- 1. 创建目录
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (0, 1, '悠然示例', 'UranExampleDir', '/uran-example', '', '', 'ri:file-list-line', 0, 0, '', '', '', 0, 0, 0, '', '', 0, 0, 100, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

SET @parentId = LAST_INSERT_ID();

-- 2. 创建页面菜单
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@parentId, 2, '悠然示例列表', 'UranExample', 'uran-example', '/uran-example/index', 'uran_example', '', 0, 1, '', '', '["GET /admin/uran-example/list"]', 0, 0, 0, '', '', 0, 0, 1, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

SET @pageId = LAST_INSERT_ID();

-- 3. 创建按钮权限（根据选项按需生成）
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@pageId, 3, '查看悠然示例', 'UranExampleView', '', '', '', '', 0, 0, '', '', '["GET /admin/uran-example/view"]', 0, 0, 0, '', '', 0, 0, 1, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@pageId, 3, '新增悠然示例', 'UranExampleAdd', '', '', '', '', 0, 0, '', '', '["POST /admin/uran-example/edit"]', 0, 0, 0, '', '', 0, 0, 2, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@pageId, 3, '编辑悠然示例', 'UranExampleEdit', '', '', '', '', 0, 0, '', '', '["POST /admin/uran-example/edit","GET /admin/uran-example/view"]', 0, 0, 0, '', '', 0, 0, 3, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@pageId, 3, '删除悠然示例', 'UranExampleDelete', '', '', '', '', 0, 0, '', '', '["POST /admin/uran-example/delete"]', 0, 0, 0, '', '', 0, 0, 4, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@pageId, 3, '导出悠然示例', 'UranExampleExport', '', '', '', '', 0, 0, '', '', '["GET /admin/uran-example/export"]', 0, 0, 0, '', '', 0, 0, 5, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
