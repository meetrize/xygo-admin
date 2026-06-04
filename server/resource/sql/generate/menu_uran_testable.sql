-- 测试记帐 菜单 SQL
-- 上级菜单ID: 0
-- ======= 顶级模式：创建目录(type=1) + 页面(type=2) + 按钮(type=3) =======

-- 1. 创建目录
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (0, 1, '测试记帐', 'UranTestableDir', '/uran-testable', '', '', 'ri:file-list-line', 0, 0, '', '', '', 0, 0, 0, '', '', 0, 0, 100, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

SET @parentId = LAST_INSERT_ID();

-- 2. 创建页面菜单
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@parentId, 2, '测试记帐列表', 'UranTestable', 'uran-testable', '/uran-testable/index', 'uran_testable', '', 0, 1, '', '', '["GET /admin/uran-testable/list"]', 0, 0, 0, '', '', 0, 0, 1, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

SET @pageId = LAST_INSERT_ID();

-- 3. 创建按钮权限（根据选项按需生成）
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@pageId, 3, '查看测试记帐', 'UranTestableView', '', '', '', '', 0, 0, '', '', '["GET /admin/uran-testable/view"]', 0, 0, 0, '', '', 0, 0, 1, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@pageId, 3, '新增测试记帐', 'UranTestableAdd', '', '', '', '', 0, 0, '', '', '["POST /admin/uran-testable/edit"]', 0, 0, 0, '', '', 0, 0, 2, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@pageId, 3, '编辑测试记帐', 'UranTestableEdit', '', '', '', '', 0, 0, '', '', '["POST /admin/uran-testable/edit","GET /admin/uran-testable/view"]', 0, 0, 0, '', '', 0, 0, 3, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@pageId, 3, '删除测试记帐', 'UranTestableDelete', '', '', '', '', 0, 0, '', '', '["POST /admin/uran-testable/delete"]', 0, 0, 0, '', '', 0, 0, 4, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@pageId, 3, '导出测试记帐', 'UranTestableExport', '', '', '', '', 0, 0, '', '', '["GET /admin/uran-testable/export"]', 0, 0, 0, '', '', 0, 0, 5, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
