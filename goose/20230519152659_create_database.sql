-- +goose Up
-- +goose StatementBegin
CREATE TABLE "chat"(
"ID" serial primary key,
"moderated" smallint not null,
"name" varchar(255),
"tg" bigint not null unique);

CREATE TABLE "user"(
"admin" smallint not null,
"ID" serial primary key,
"name" varchar(255),
"tg" bigint not null);

CREATE TABLE "ban"(
"ID" serial primary key,
"warning" int,
"banstart" time,
"banend" time,
"banreason" varchar(200),
"tg" int);

create table "bannedwords"(
    "id" serial primary key,
    "word" varchar(60) not null
)

create table "okurl"
(
    "id" serial primary key,
    "url" varchar(300) not null
)

INSERT INTO "bannedwords" ("id", "word") VALUES
    (1315, 'spam'),
    (2, 'аборт'),
    (1, 'анус'),
    (4, 'беспезды'),
    (3, 'бздун'),
    (5, 'бздюх'),
    (7, 'блудилище'),
    (6, 'бля'),
    (26, 'бляд'),
    (8, 'блядво'),
    (9, 'блядеха'),
    (25, 'бляди'),
    (10, 'блядина'),
    (11, 'блядистка'),
    (12, 'блядище'),
    (13, 'блядки'),
    (14, 'блядование'),
    (15, 'блядовать'),
    (16, 'блядовитый'),
    (17, 'блядовозка'),
    (18, 'блядолиз'),
    (19, 'блядоход'),
    (20, 'блядский'),
    (21, 'блядство'),
    (22, 'блядствовать'),
    (23, 'блядун'),
    (24, 'блядь'),
    (27, 'блядюга'),
    (28, 'блядюра'),
    (29, 'блядюшка'),
    (30, 'блядюшник'),
    (31, 'бордель'),
    (32, 'вагина'),
    (33, 'вафлист'),
    (34, 'вжопить'),
    (35, 'вжопиться'),
    (36, 'вздрачивание'),
    (37, 'вздрачивать'),
    (38, 'вздрачиваться'),
    (39, 'вздрочить'),
    (40, 'вздрочиться'),
    (42, 'вздрючивание'),
    (43, 'вздрючивать'),
    (41, 'вздрючить'),
    (44, 'взъебка'),
    (46, 'взъебнуть'),
    (45, 'взъебщик'),
    (47, 'вислозадая'),
    (48, 'влагалище'),
    (49, 'вхуйнуть'),
    (50, 'вхуйнуться'),
    (53, 'вхуя'),
    (51, 'вхуякать'),
    (52, 'вхуякаться'),
    (54, 'вхуякивать'),
    (55, 'вхуякиваться'),
    (56, 'вхуякнуть'),
    (57, 'вхуякнуться'),
    (58, 'вхуяривание'),
    (59, 'вхуяривать'),
    (60, 'вхуяриваться'),
    (61, 'вхуярить'),
    (62, 'вхуяриться'),
    (63, 'вхуячивание'),
    (64, 'вхуячивать'),
    (65, 'вхуячиваться'),
    (66, 'вхуячить'),
    (67, 'вхуячиться'),
    (68, 'вхуяшивать'),
    (69, 'вхуяшиваться'),
    (70, 'вхуяшить'),
    (71, 'вхуяшиться'),
    (72, 'въебать'),
    (73, 'въебаться'),
    (74, 'въебашивать'),
    (75, 'въебашиваться'),
    (76, 'въебашить'),
    (77, 'въебашиться'),
    (78, 'въебенивать'),
    (79, 'въебениваться'),
    (80, 'въебенить'),
    (81, 'въебениться'),
    (82, 'выблядок'),
    (83, 'выебанный'),
    (84, 'выебат'),
    (205, 'выебать'),
    (85, 'выебаться'),
    (90, 'высераться'),
    (86, 'высрать'),
    (87, 'высраться'),
    (88, 'выссать'),
    (89, 'выссаться'),
    (91, 'выссереть'),
    (92, 'говнецо'),
    (93, 'говнистый'),
    (94, 'говниться'),
    (95, 'говно'),
    (96, 'говновоз'),
    (97, 'говнодав'),
    (98, 'говноеб'),
    (99, 'говноед'),
    (101, 'говномер'),
    (100, 'говномес'),
    (102, 'говносерка'),
    (103, 'говнюк'),
    (104, 'голожопая'),
    (105, 'гомик'),
    (106, 'гомосек'),
    (107, 'гондон'),
    (108, 'гонорея'),
    (109, 'давалка'),
    (110, 'двужопник'),
    (111, 'дерьмо'),
    (113, 'дерьмовый'),
    (112, 'дерьмоед'),
    (114, 'дилдо'),
    (115, 'додрочить'),
    (116, 'додрочиться'),
    (117, 'доебать'),
    (118, 'доебаться'),
    (119, 'доебенивать'),
    (120, 'доебениваться'),
    (121, 'доебенить'),
    (122, 'доебениться'),
    (123, 'долбоеб'),
    (124, 'допиздить'),
    (125, 'допиздиться'),
    (170, 'допиздоболивать'),
    (171, 'допиздоболиваться'),
    (172, 'допиздоболиться'),
    (126, 'допиздовать'),
    (127, 'допиздоваться'),
    (128, 'допиздовывать'),
    (129, 'допиздовываться'),
    (130, 'допиздохать'),
    (131, 'допиздохаться'),
    (132, 'допиздохивать'),
    (133, 'допиздохиваться'),
    (136, 'допиздошивать'),
    (137, 'допиздошиваться'),
    (134, 'допиздошить'),
    (135, 'допиздошиться'),
    (173, 'допиздюкать'),
    (174, 'допиздюкаться'),
    (175, 'допиздюкивать'),
    (176, 'допиздюкиваться'),
    (140, 'допиздюливать'),
    (141, 'допиздюливаться'),
    (138, 'допиздюлить'),
    (139, 'допиздюлиться'),
    (144, 'допиздюривать'),
    (145, 'допиздюриваться'),
    (142, 'допиздюрить'),
    (143, 'допиздюриться'),
    (146, 'допиздюхать'),
    (147, 'допиздюхаться'),
    (148, 'допиздюхивать'),
    (149, 'допиздюхиваться'),
    (150, 'допиздякать'),
    (151, 'допиздякаться'),
    (152, 'допиздякивать'),
    (153, 'допиздякиваться'),
    (156, 'допиздяривать'),
    (157, 'допиздяриваться'),
    (154, 'допиздярить'),
    (155, 'допиздяриться'),
    (158, 'допиздяхать'),
    (159, 'допиздяхаться'),
    (160, 'допиздяхивать'),
    (161, 'допиздяхиваться'),
    (164, 'допиздячивать'),
    (165, 'допиздячиваться'),
    (162, 'допиздячить'),
    (163, 'допиздячиться'),
    (168, 'допиздяшивать'),
    (169, 'допиздяшиваться'),
    (166, 'допиздяшить'),
    (167, 'допиздяшиться'),
    (177, 'допизживать'),
    (178, 'дотрахать'),
    (179, 'дотрахаться'),
    (180, 'дохуйнуть'),
    (181, 'дохуякать'),
    (182, 'дохуякаться'),
    (183, 'дохуякивать'),
    (184, 'дохуякиваться'),
    (185, 'дохуяривать'),
    (186, 'дохуяриваться'),
    (187, 'дохуярить'),
    (188, 'дохуяриться'),
    (191, 'дохуячивать'),
    (192, 'дохуячиваться'),
    (189, 'дохуячить'),
    (190, 'дохуячиться'),
    (193, 'дрисня'),
    (194, 'дристать'),
    (195, 'дристун'),
    (196, 'дроченье'),
    (197, 'дрочилыцик'),
    (198, 'дрочить'),
    (199, 'дрочиться'),
    (200, 'дрочка'),
    (201, 'дрючить'),
    (202, 'дрючиться'),
    (203, 'дурак'),
    (204, 'дуроеб'),
    (206, 'ебало'),
    (207, 'ебальник'),
    (208, 'ебальные'),
    (209, 'ебальный'),
    (210, 'ебанатик'),
    (211, 'ебанашка'),
    (212, 'ебанутый'),
    (213, 'ебануть'),
    (214, 'ебануться'),
    (216, 'ебат'),
    (215, 'ебать'),
    (218, 'ебатьс'),
    (217, 'ебаться'),
    (219, 'ебитесь'),
    (220, 'ебло'),
    (221, 'еблом'),
    (222, 'еблысь'),
    (223, 'ебля'),
    (224, 'ебнуть'),
    (225, 'ебнуться'),
    (226, 'ебня'),
    (227, 'ебучий'),
    (232, 'жирнозадый'),
    (233, 'жопа'),
    (235, 'жопастая'),
    (237, 'жопенци'),
    (238, 'жопища'),
    (239, 'жопка'),
    (240, 'жопник'),
    (236, 'жопоеб'),
    (234, 'жопой'),
    (241, 'жополиз'),
    (242, 'жополизание'),
    (243, 'жопоногий'),
    (244, 'жопочка'),
    (245, 'жопочник'),
    (246, 'жопство'),
    (247, 'жопу'),
    (248, 'забздеть'),
    (249, 'заблядовать'),
    (250, 'заблядоваться'),
    (251, 'задница'),
    (252, 'задрачивать'),
    (253, 'задрачиваться'),
    (254, 'задроченный'),
    (255, 'задрочить'),
    (256, 'задрочиться'),
    (257, 'задрючить'),
    (258, 'задрючиться'),
    (259, 'заебанный'),
    (262, 'заебательская'),
    (260, 'заебать'),
    (261, 'заебаться'),
    (263, 'заебашивать'),
    (264, 'заебашиваться'),
    (265, 'заебашить'),
    (266, 'заебашиться'),
    (267, 'заебенивать'),
    (268, 'заебениваться'),
    (269, 'заебенить'),
    (270, 'заебениться'),
    (228, 'заебла'),
    (271, 'залупа'),
    (273, 'залупаться'),
    (274, 'залупенить'),
    (275, 'залупень'),
    (278, 'залупистый'),
    (276, 'залупить'),
    (277, 'залупляться'),
    (272, 'залупу'),
    (279, 'запиздарить'),
    (281, 'запизденелый'),
    (280, 'запизденная'),
    (282, 'запиздить'),
    (283, 'запиздиться'),
    (284, 'запиздоболивать'),
    (285, 'запиздоболиваться'),
    (286, 'запиздоболить'),
    (287, 'запиздоболиться'),
    (288, 'запиздовать'),
    (289, 'запиздоваться'),
    (290, 'запиздовывать'),
    (291, 'запиздовываться'),
    (292, 'запиздохать'),
    (295, 'запиздошивать'),
    (296, 'запиздошиваться'),
    (293, 'запиздошить'),
    (294, 'запиздошиться'),
    (297, 'запиздюкать'),
    (298, 'запиздюкаться'),
    (299, 'запиздюкивать'),
    (300, 'запиздюкиваться'),
    (303, 'запиздюливать'),
    (304, 'запиздюливаться'),
    (301, 'запиздюлить'),
    (302, 'запиздюлиться'),
    (307, 'запиздюривать'),
    (308, 'запиздюриваться'),
    (305, 'запиздюрить'),
    (306, 'запиздюриться'),
    (309, 'запиздюхать'),
    (310, 'запиздюхаться'),
    (311, 'запиздюхивать'),
    (312, 'запиздюхиваться'),
    (315, 'запиздючивать'),
    (316, 'запиздючиваться'),
    (313, 'запиздючить'),
    (314, 'запиздючиться'),
    (317, 'засранец'),
    (318, 'засранка'),
    (319, 'засранный'),
    (320, 'засратый'),
    (321, 'засрать'),
    (322, 'засраться'),
    (323, 'зассать'),
    (324, 'затраханный'),
    (325, 'затрахать'),
    (326, 'затрахаться'),
    (327, 'затрахивать'),
    (328, 'затрахиваться'),
    (329, 'захуить'),
    (330, 'захуйнуть'),
    (331, 'захуйнуться'),
    (332, 'захуякать'),
    (333, 'захуякаться'),
    (334, 'захуякивать'),
    (335, 'захуякиваться'),
    (338, 'захуяривать'),
    (339, 'захуяриваться'),
    (336, 'захуярить'),
    (337, 'захуяриться'),
    (342, 'захуячивать'),
    (343, 'захуячиваться'),
    (340, 'захуячить'),
    (341, 'захуячиться'),
    (346, 'захуяшивать'),
    (347, 'захуяшиваться'),
    (344, 'захуяшить'),
    (345, 'захуяшиться'),
    (348, 'злоебучий'),
    (349, 'издрочиться'),
    (350, 'измандить'),
    (351, 'измандиться'),
    (352, 'измандовать'),
    (353, 'измандоваться'),
    (354, 'измандовывать'),
    (355, 'измандовываться'),
    (366, 'изъеб'),
    (356, 'изъебать'),
    (357, 'изъебаться'),
    (360, 'изъебашивать'),
    (361, 'изъебашиваться'),
    (358, 'изъебашить'),
    (359, 'изъебашиться'),
    (364, 'изъебенивать'),
    (365, 'изъебениваться'),
    (362, 'изъебенить'),
    (363, 'изъебениться'),
    (367, 'испиздеться'),
    (368, 'испиздить'),
    (369, 'испражнение'),
    (370, 'испражняться'),
    (371, 'исхуякать'),
    (372, 'исхуякаться'),
    (373, 'исхуякивать'),
    (374, 'исхуякиваться'),
    (377, 'исхуяривать'),
    (375, 'исхуярить'),
    (376, 'исхуяриться'),
    (378, 'какать'),
    (379, 'какашка'),
    (380, 'кастрат'),
    (381, 'кастрировать'),
    (382, 'клитор'),
    (383, 'клоака'),
    (384, 'кнахт'),
    (385, 'кончить'),
    (386, 'косоебить'),
    (387, 'косоебиться'),
    (388, 'кривохуй'),
    (389, 'курва'),
    (390, 'курвиный'),
    (391, 'лахудра'),
    (392, 'лох'),
    (394, 'лохматка'),
    (393, 'лохудра'),
    (395, 'манда'),
    (397, 'мандавоха'),
    (396, 'мандавошка'),
    (398, 'мандить'),
    (399, 'мандиться'),
    (400, 'мандоватая'),
    (401, 'мандовать'),
    (402, 'мандохать'),
    (403, 'мандохаться'),
    (404, 'мандохивать'),
    (405, 'мандохиваться'),
    (406, 'мандошить'),
    (407, 'мастурбатор'),
    (408, 'минет'),
    (409, 'минетить'),
    (410, 'минетка'),
    (411, 'минетчик'),
    (412, 'минетчица'),
    (413, 'мозгоеб'),
    (414, 'мозгоебатель'),
    (415, 'мозгоебать'),
    (416, 'мозгоебка'),
    (417, 'мокрожопый'),
    (418, 'мокропиздая'),
    (419, 'моча'),
    (420, 'мочиться'),
    (421, 'мудак'),
    (422, 'мудашвили'),
    (1316, 'мудила'),
    (423, 'мудило'),
    (424, 'мудильщик'),
    (425, 'мудистый'),
    (426, 'мудить'),
    (427, 'мудоеб'),
    (434, 'набздеть'),
    (435, 'наблядоваться'),
    (440, 'надристать'),
    (229, 'надроченный'),
    (436, 'надроченный'),
    (437, 'надрочивать'),
    (438, 'надрочить'),
    (439, 'надрочиться'),
    (428, 'наебанный'),
    (441, 'наебать'),
    (442, 'наебаться'),
    (429, 'наебка'),
    (443, 'наебнуть'),
    (444, 'наебнуться'),
    (430, 'наебщик'),
    (431, 'наебывать'),
    (432, 'наебываться'),
    (433, 'наебыш'),
    (445, 'накакать'),
    (446, 'накакаться'),
    (447, 'накакивать'),
    (448, 'напиздить'),
    (449, 'напиздошить'),
    (450, 'напиздюрить'),
    (451, 'напиздюриться'),
    (452, 'насрать'),
    (453, 'насраться'),
    (454, 'нассать'),
    (455, 'нассаться'),
    (456, 'натрахать'),
    (457, 'натрахаться'),
    (458, 'натрахивать'),
    (459, 'натрахиваться'),
    (460, 'нахуякать'),
    (461, 'нахуякаться'),
    (462, 'нахуякивать'),
    (463, 'нахуякиваться'),
    (467, 'нахуяривать'),
    (468, 'нахуяриваться'),
    (464, 'нахуярить'),
    (465, 'нахуяриться'),
    (466, 'нахуяриться'),
    (471, 'нахуячивать'),
    (472, 'нахуячиваться'),
    (469, 'нахуячить'),
    (470, 'нахуячиться'),
    (473, 'нахуяшить'),
    (474, 'недоебанный'),
    (475, 'недоносок'),
    (476, 'неебущий'),
    (477, 'нищеебство'),
    (479, 'обдристанный'),
    (480, 'обдристать'),
    (481, 'обдрочиться'),
    (482, 'обосранец'),
    (483, 'обосранная'),
    (484, 'обосраный'),
    (485, 'обосрать'),
    (486, 'обосраться'),
    (487, 'обоссанец'),
    (488, 'обоссаный'),
    (489, 'обоссать'),
    (490, 'обоссаться'),
    (491, 'обоссаться'),
    (492, 'обоссывать'),
    (493, 'обоссываться'),
    (494, 'обпиздить'),
    (495, 'обпиздиться'),
    (496, 'обпиздовать'),
    (497, 'обпиздоваться'),
    (498, 'обпиздовывать'),
    (499, 'обпиздовываться'),
    (500, 'обпиздохать'),
    (501, 'обпиздохаться'),
    )
INSERT INTO "bannedwords" ("id", "word") VALUES(
    (502, 'обпиздохивать'),
    (503, 'обпиздохиваться'),
    (504, 'обпиздошить'),
    (505, 'обтрахать'),
    (506, 'обтрахаться'),
    (507, 'обтрахивать'),
    (508, 'обтрахиваться'),
    (509, 'обхуярить'),
    (510, 'обхуяриться'),
    (511, 'обхуячить'),
    (512, 'объебать'),
    (513, 'объебаться'),
    (514, 'объебенить'),
    (230, 'объебешь'),
    (515, 'объебнуть'),
    (516, 'объебон'),
    (517, 'одинхуй'),
    (518, 'однапизда'),
    (519, 'однохуйственно'),
    (520, 'оебать'),
    (521, 'оебашивать'),
    (522, 'оебашить'),
    (523, 'оебенивать'),
    (524, 'оебенить'),
    (478, 'оебыват'),
    (525, 'опедерастить'),
    (526, 'опизденеть'),
    (528, 'опизденно'),
    (527, 'опизденный'),
    (529, 'опиздеть'),
    (530, 'опиздить'),
    (533, 'остоебенило'),
    (532, 'остоебенить'),
    (531, 'остоебеть'),
    (535, 'остопиздело'),
    (534, 'остопиздеть'),
    (536, 'остохуело'),
    (537, 'остохуеть'),
    (538, 'отдрачивать'),
    (539, 'отдрачиваться'),
    (540, 'отдрочить'),
    (541, 'отдрочиться'),
    (542, 'отпиздить'),
    (543, 'отпиздошить'),
    (546, 'отпиздяшивание'),
    (547, 'отпиздяшивать'),
    (548, 'отпиздяшиваться'),
    (544, 'отпиздяшить'),
    (545, 'отпиздяшиться'),
    (549, 'отсасывать'),
    (550, 'отсасываться'),
    (551, 'отсосать'),
    (552, 'отсосаться'),
    (553, 'оттраханная'),
    (554, 'оттрахать'),
    (555, 'оттрахаться'),
    (556, 'оттрахивать'),
    (557, 'оттрахиваться'),
    (558, 'отхерачить'),
    (559, 'отхуякать'),
    (560, 'отхуякаться'),
    (561, 'отхуякивать'),
    (562, 'отхуякиваться'),
    (565, 'отхуяривать'),
    (566, 'отхуяриваться'),
    (563, 'отхуярить'),
    (564, 'отхуяриться'),
    (569, 'отхуячивать'),
    (570, 'отхуячиваться'),
    (567, 'отхуячить'),
    (568, 'отхуячиться'),
    (573, 'отхуяшивать'),
    (574, 'отхуяшиваться'),
    (571, 'отхуяшить'),
    (572, 'отхуяшиться'),
    (575, 'отъебать'),
    (580, 'отъебашивание'),
    (581, 'отъебашивать'),
    (582, 'отъебашиваться'),
    (579, 'отъебашить'),
    (585, 'отъебенивать'),
    (586, 'отъебениваться'),
    (583, 'отъебенить'),
    (584, 'отъебениться'),
    (1312, 'отъебись'),
    (587, 'отъебнуть'),
    (576, 'отъебывание'),
    (577, 'отъебывать'),
    (578, 'отъебываться'),
    (588, 'отьебаться'),
    (589, 'отьебашиться'),
    (590, 'отьебенивание'),
    (1313, 'отьебись'),
    (591, 'отьебнуться'),
    (592, 'охуевать'),
    (593, 'охуевающий'),
    (594, 'охуевший'),
    (595, 'охуение'),
    (596, 'охуенно'),
    (597, 'охуенные'),
    (598, 'охуеть'),
    (599, 'охуительно'),
    (600, 'охуительный'),
    (601, 'охуякать'),
    (602, 'охуякаться'),
    (603, 'охуякивать'),
    (604, 'охуякиваться'),
    (605, 'охуякнуть'),
    (606, 'охуякнуться'),
    (609, 'охуяривать'),
    (610, 'охуяриваться'),
    (607, 'охуярить'),
    (608, 'охуяриться'),
    (613, 'охуячивать'),
    (614, 'охуячиваться'),
    (611, 'охуячить'),
    (612, 'охуячиться'),
    (617, 'охуяшивать'),
    (618, 'охуяшиваться'),
    (615, 'охуяшить'),
    (616, 'охуяшиться'),
    (619, 'очко'),
    (621, 'падла'),
    (622, 'падлюка'),
    (623, 'педераст'),
    (624, 'педерастина'),
    (625, 'педерастический'),
    (626, 'педерастия'),
    (627, 'педик'),
    (628, 'педрило'),
    (629, 'пежить'),
    (630, 'пенис'),
    (631, 'пердеж'),
    (632, 'пердеть'),
    (620, 'перднуть'),
    (633, 'пердун'),
    (634, 'перебздеть'),
    (635, 'передрачивать'),
    (636, 'передрочить'),
    (637, 'передрочиться'),
    (638, 'переебаться'),
    (639, 'переебашить'),
    (640, 'перетрахать'),
    (641, 'перетрахаться'),
    (642, 'перетрахивать'),
    (643, 'перетрахиваться'),
    (644, 'перехуйнуть'),
    (645, 'перехуйнуться'),
    (648, 'перехуякать'),
    (649, 'перехуякаться'),
    (650, 'перехуякивать'),
    (651, 'перехуякиваться'),
    (646, 'перехуякнуть'),
    (647, 'перехуякнуться'),
    (654, 'перехуяривать'),
    (655, 'перехуяриваться'),
    (652, 'перехуярить'),
    (653, 'перехуяриться'),
    (658, 'перехуячивать'),
    (656, 'перехуячить'),
    (657, 'перехуячиться'),
    (660, 'пидор'),
    (659, 'пидорас'),
    (661, 'пизда'),
    (662, 'пизданутая'),
    (663, 'пиздануть'),
    (664, 'пиздануться'),
    (665, 'пиздато'),
    (666, 'пизденка'),
    (667, 'пизденочка'),
    (669, 'пизденыш'),
    (668, 'пиздень'),
    (670, 'пиздеть'),
    (671, 'пиздец'),
    (672, 'пиздища'),
    (673, 'пиздобол'),
    (674, 'пиздовать'),
    (675, 'пиздолиз'),
    (676, 'пиздомол'),
    (677, 'пиздосос'),
    (678, 'пиздоход'),
    (679, 'пиздуй'),
    (680, 'пиздун'),
    (681, 'пиздюга'),
    (685, 'пиздюк'),
    (686, 'пиздюкать'),
    (687, 'пиздюкаться'),
    (682, 'пиздюлей'),
    (683, 'пиздюли'),
    (684, 'пиздюлина'),
    (688, 'пиздюшка'),
    (689, 'пиздякать'),
    (690, 'пиздятина'),
    (691, 'пиздятиной'),
    (692, 'пиздячий'),
    (693, 'писька'),
    (694, 'писюлек'),
    (695, 'плоскозадая'),
    (699, 'поблудить'),
    (700, 'поблядовать'),
    (701, 'поблядушка'),
    (702, 'подосрать'),
    (703, 'подосраться'),
    (704, 'подоссать'),
    (705, 'подпиздить'),
    (706, 'подпиздовать'),
    (707, 'подпиздоваться'),
    (708, 'подпиздовывать'),
    (709, 'подпиздовываться'),
    (710, 'подпиздохать'),
    (711, 'подпиздохаться'),
    (712, 'подпиздохивать'),
    (713, 'подпиздохиваться'),
    (716, 'подпиздошивать'),
    (714, 'подпиздошить'),
    (715, 'подпиздошиться'),
    (717, 'подпиздякать'),
    (718, 'подпиздякаться'),
    (719, 'подпиздякивать'),
    (720, 'подпиздякиваться'),
    (723, 'подпиздяривать'),
    (724, 'подпиздяриваться'),
    (721, 'подпиздярить'),
    (722, 'подпиздяриться'),
    (725, 'подпиздяхать'),
    (726, 'подпиздяхаться'),
    (727, 'подпиздяхивать'),
    (728, 'подпиздяхиваться'),
    (731, 'подпиздячивать'),
    (732, 'подпиздячиваться'),
    (729, 'подпиздячить'),
    (730, 'подпиздячиться'),
    (735, 'подпиздяшивать'),
    (736, 'подпиздяшиваться'),
    (733, 'подпиздяшить'),
    (734, 'подпиздяшиться'),
    (737, 'подристывать'),
    (738, 'подрочить'),
    (739, 'подсирать'),
    (742, 'подхуякать'),
    (743, 'подхуякаться'),
    (744, 'подхуякивать'),
    (745, 'подхуякиваться'),
    (740, 'подхуякнуть'),
    (741, 'подхуякнуться'),
    (748, 'подхуяривать'),
    (749, 'подхуяриваться'),
    (746, 'подхуярить'),
    (747, 'подхуяриться'),
    (750, 'подхуячивать'),
    (752, 'подхуячивать'),
    (753, 'подхуячиваться'),
    (751, 'подхуячиться'),
    (756, 'подхуяшивать'),
    (757, 'подхуяшиваться'),
    (754, 'подхуяшить'),
    (755, 'подхуяшиться'),
    (758, 'подъеб'),
    (759, 'подъебать'),
    (760, 'подъебаться'),
    (761, 'подъебашить'),
    (763, 'подъебка'),
    (762, 'подъебнуть'),
    (764, 'подъебывать'),
    (765, 'подъябывать'),
    (766, 'поебанный'),
    (231, 'поебать'),
    (767, 'поебать'),
    (768, 'поебаться'),
    (769, 'поебень'),
    (770, 'поебистика'),
    (771, 'поебон'),
    (772, 'поебончик'),
    (696, 'поебочка'),
    (697, 'поебывать'),
    (698, 'поебываться'),
    (773, 'попердеть'),
    (774, 'попердеться'),
    (775, 'попердывать'),
    (776, 'попизденная'),
    (777, 'попиздеть'),
    (778, 'попиздистее'),
    (779, 'попиздить'),
    (780, 'попиздиться'),
    (782, 'попиздоболивать'),
    (783, 'попиздоболиваться'),
    (784, 'попиздоболить'),
    (785, 'попиздоболиться'),
    (781, 'попиздоватей'),
    (786, 'попиздовать'),
    (787, 'попиздоваться'),
    (788, 'попиздовывать'),
    (789, 'попиздовываться'),
    (790, 'попиздохать'),
    (791, 'попиздохаться'),
    (792, 'попиздохивать'),
    (793, 'попиздохиваться'),
    (796, 'попиздошивать'),
    (797, 'попиздошиваться'),
    (794, 'попиздошить'),
    (795, 'попиздошиться'),
    (798, 'попиздюкать'),
    (799, 'попиздюкаться'),
    (800, 'попиздюкивать'),
    (801, 'попиздюкиваться'),
    (804, 'попиздюливать'),
    (805, 'попиздюливаться'),
    (802, 'попиздюлить'),
    (803, 'попиздюлиться'),
    (808, 'попиздюривать'),
    (809, 'попиздюриваться'),
    (806, 'попиздюрить'),
    (807, 'попиздюриться'),
    (810, 'попиздюхать'),
    (811, 'попиздюхаться'),
    (812, 'попиздюхивать'),
    (813, 'попиздюхиваться'),
    (814, 'попиздякать'),
    (815, 'попиздякаться'),
    (816, 'попиздякивать'),
    (817, 'попиздякиваться'),
    (820, 'попиздяривать'),
    (821, 'попиздяриваться'),
    (818, 'попиздярить'),
    (819, 'попиздяриться'),
    (822, 'попиздяхать'),
    (823, 'попиздяхаться'),
    (824, 'попиздяхивать'),
    (825, 'попиздяхиваться'),
    (828, 'попиздячивать'),
    (829, 'попиздячиваться'),
    (826, 'попиздячить'),
    (827, 'попиздячиться'),
    (832, 'попиздяшивать'),
    (833, 'попиздяшиваться'),
    (830, 'попиздяшить'),
    (831, 'попиздяшиться'),
    (834, 'попизживать'),
    (835, 'попизживаться'),
    (836, 'потаскун'),
    (837, 'потаскуха'),
    (838, 'потраханная'),
    (839, 'потрахать'),
    (840, 'потрахаться'),
    (841, 'потрахивать'),
    (842, 'потрахиваться'),
    (843, 'похер'),
    (844, 'похуист'),
    (845, 'похуякать'),
    (846, 'похуякаться'),
    (847, 'похуякивать'),
    (848, 'похуякиваться'),
    (851, 'похуяривать'),
    (852, 'похуяриваться'),
    (849, 'похуярить'),
    (850, 'похуяриться'),
    (855, 'похуячивать'),
    (856, 'похуячиваться'),
    (853, 'похуячить'),
    (854, 'похуячиться'),
    (859, 'похуяшивать'),
    (860, 'похуяшиваться'),
    (857, 'похуяшить'),
    (858, 'похуяшиться'),
    (861, 'поц'),
    (862, 'пошмариться'),
    (863, 'поябывать'),
    (864, 'приебать'),
    (865, 'приебаться'),
    (870, 'приебашивать'),
    (871, 'приебашиваться'),
    (868, 'приебашить'),
    (869, 'приебашиться'),
    (874, 'приебенивать'),
    (875, 'приебениваться'),
    (872, 'приебенить'),
    (873, 'приебениться'),
    (876, 'приебехать'),
    (877, 'приебехаться'),
    (878, 'приебехивать'),
    (879, 'приебехиваться'),
    (880, 'приебистый'),
    (883, 'приебуривать'),
    (884, 'приебуриваться'),
    (881, 'приебурить'),
    (882, 'приебуриться'),
    (866, 'приебывать'),
    (867, 'приебываться'),
    (885, 'прижопить'),
    (886, 'прижопывать'),
    (887, 'прикинуть'),
    (890, 'примавдовывать'),
    (904, 'примандехать'),
    (905, 'примандехаться'),
    (906, 'примандехивать'),
    (907, 'примандехиваться'),
    (940, 'примандить'),
    (941, 'примандиться'),
    (888, 'примандовать'),
    (889, 'примандоваться'),
    (891, 'примандовываться'),
    (892, 'примандохать'),
    (893, 'примандохаться'),
    (894, 'примандохивать'),
    (895, 'примандохиваться'),
    (898, 'примандошивать'),
    (899, 'примандошиваться'),
    (896, 'примандошить'),
    (897, 'примандошиться'),
    (900, 'примандюкать'),
    (901, 'примандюкаться'),
    (902, 'примандюкивать'),
    (903, 'примандюкиваться'),
    (910, 'примандюливать'),
    (911, 'примандюливаться'),
    (908, 'примандюлить'),
    (909, 'примандюлиться'),
    (914, 'примандюривать'),
    (915, 'примандюриваться'),
    (912, 'примандюрить'),
    (913, 'примандюриться'),
    (916, 'примандякать'),
    (917, 'примандякаться'),
    (918, 'примандякивать'),
    (919, 'примандякиваться'),
    (922, 'примандяривать'),
    (923, 'примандяриваться'),
    (920, 'примандярить'),
    (921, 'примандяриться'),
    (924, 'примандяхать'),
    (925, 'примандяхаться'),
    (926, 'примандяхивать'),
    (927, 'примандяхиваться'),
    (930, 'примандячивать'),
    (931, 'примандячиваться'),
    (928, 'примандячить'),
    (929, 'примандячиться'),
    (934, 'примандяшивать'),
    (935, 'примандяшиваться'),
    (932, 'примандяшить'),
    (933, 'примандяшиться'),
    (936, 'примудохать'),
    (937, 'примудохаться'),
    (938, 'примудохивать'),
    (939, 'примудохиваться'),
    (942, 'припизденный'),
    (943, 'припиздень'),
    (944, 'припиздить'),
    (945, 'припиздиться'),
    (948, 'припиздовать'),
    (949, 'припиздоваться'),
    (950, 'припиздовывать'),
    (951, 'припиздовываться'),
    (952, 'припиздохать'),
    (953, 'припиздохаться'),
    (954, 'припиздохивать'),
    (955, 'припиздохиваться'),
    (958, 'припиздошивать'),
    (959, 'припиздошиваться'),
    (956, 'припиздошить'),
    (957, 'припиздошиться'),
    (998, 'припиздронивать'),
    (999, 'припиздрониваться'),
    (996, 'припиздронить'),
    (997, 'припиздрониться'),
    (946, 'припиздывать'),
    (947, 'припиздываться'),
    (960, 'припиздюкать'),
    (961, 'припиздюкаться'),
    (962, 'припиздюкивать'),
    (963, 'припиздюкиваться'),
    (966, 'припиздюливать'),
    (967, 'припиздюливаться'),
    (964, 'припиздюлить'),
    (965, 'припиздюлиться'),
    (970, 'припиздюривать'),
    (972, 'припиздюриваться'),
    (968, 'припиздюрить'),
    (969, 'припиздюриться'),
    (971, 'припиздюхать'),
    (973, 'припиздюхаться'),
    (974, 'припиздюхивать'),
    (975, 'припиздюхиваться'),
    (976, 'припиздякать'),
    (977, 'припиздякаться'),
    (978, 'припиздякивать'),
    (979, 'припиздякиваться'),
    (982, 'припиздяривать'),
    (983, 'припиздяриваться'),
    (980, 'припиздярить'),
    (981, 'припиздяриться'),
    (984, 'припиздяхать'),
    (985, 'припиздяхаться'),
    (986, 'припиздяхивать'),
    (987, 'припиздяхиваться'),
    (990, 'припиздячивать'),
    (991, 'припиздячиваться'),
    (988, 'припиздячить'),
    (989, 'припиздячиться'),
    (994, 'припиздяшивать'),
    (995, 'припиздяшиваться'),
    (992, 'припиздяшить'),
    (993, 'припиздяшиться'),
    (1000, 'припизживать'),
    (1001, 'припизживаться'),
    (1019, 'притрахаться'),
    (1002, 'прихуеть'),
    (1003, 'прихуякать'),
    (1004, 'прихуякаться'),
    (1005, 'прихуякивать'),
    (1006, 'прихуякиваться'),
    (1009, 'прихуяривать'),
    (1010, 'прихуяриваться'),
    (1007, 'прихуярить'),
    (1008, 'прихуяриться'),
    (1013, 'прихуячивать'),
    (1014, 'прихуячиваться'),
    (1011, 'прихуячить'),
    (1012, 'прихуячиться'),
    (1017, 'прихуяшивать'),
    (1018, 'прихуяшиваться'),
    (1015, 'прихуяшить'),
    (1016, 'прихуяшиться'),
    (1020, 'проблядовать'),
    (1022, 'проблядушка'),
    (1021, 'проблядь'),
    (1023, 'продрачивать'),
    (1024, 'продрачиваться'),
    (1025, 'продрочить'),
    (1026, 'продрочиться'),
    (1027, 'проебать'),
    (1028, 'проебаться'),
    (1031, 'проебашивать'),
    (1035, 'проебашивать'),
    (1032, 'проебашиваться'),
    (1036, 'проебашиваться'),
    (1029, 'проебашить'),
    (1030, 'проебашиться'),
    (1033, 'проебенить'),
    (1034, 'проебениться'),
    (1037, 'проебывать'),
    (1038, 'проебываться'),
    (1039, 'пропиздить'),
    (1040, 'пропиздиться'),
    (1041, 'пропиздоболивать'),
    (1042, 'пропиздоболиваться'),
    (1043, 'пропиздоболить'),
    (1044, 'пропиздоболиться'),
    (1045, 'пропиздовать'),
    (1046, 'пропиздоваться'),
    (1047, 'пропиздовывать'),
    (1048, 'пропиздовываться'),
    (1094, 'пропиздон'),
    (1049, 'пропиздохать'),
    (1050, 'пропиздохаться'),
    (1051, 'пропиздохивать'),
    (1052, 'пропиздохиваться'),
    (1055, 'пропиздошивать'),
    (1056, 'пропиздошиваться'),
    (1053, 'пропиздошить'),
    (1054, 'пропиздошиться'),
    (1057, 'пропиздюкать'),
    (1058, 'пропиздюкаться'),
    (1059, 'пропиздюкивать'),
    (1060, 'пропиздюкиваться'),
    (1063, 'пропиздюливать'),
    (1064, 'пропиздюливаться'),
    (1061, 'пропиздюлить'),
    (1062, 'пропиздюлиться'),
    (1067, 'пропиздюривать'),
    (1068, 'пропиздюриваться'),
    (1065, 'пропиздюрить'),
    (1066, 'пропиздюриться'),
    (1069, 'пропиздюхать'),
    (1070, 'пропиздюхаться'),
    (1071, 'пропиздюхивать'),
    (1072, 'пропиздюхиваться'),
    (1073, 'пропиздякать'),
    (1074, 'пропиздякаться'),
    (1075, 'пропиздякивать'),
    (1076, 'пропиздякиваться'),
    (1079, 'пропиздяривать'),
    (1080, 'пропиздяриваться'),
    (1077, 'пропиздярить'),
    (1078, 'пропиздяриться'),
    (1081, 'пропиздяхать'),
    (1082, 'пропиздяхивать'),
    (1083, 'пропиздяхиваться'),
    (1086, 'пропиздячивать'),
    (1087, 'пропиздячиваться'),
    (1084, 'пропиздячить'),
    (1085, 'пропиздячиться'),
    (1090, 'пропиздяшивать'),
    (1091, 'пропиздяшиваться'),
    (1088, 'пропиздяшить'),
    (1089, 'пропиздяшиться'),
    (1092, 'пропизживать'),
    (1093, 'пропизживаться'),
    (1095, 'прохуякать'),
    (1096, 'прохуякаться'),
    (1097, 'прохуякивать'),
    (1098, 'прохуякиваться'),
    (1101, 'прохуяривать'),
    (1102, 'прохуяриваться'),
    (1099, 'прохуярить'),
    (1100, 'прохуяриться'),
    (1105, 'прохуячивать'),
    (1106, 'прохуячиваться'),
    (1103, 'прохуячить'),
    (1104, 'прохуячиться'),
    (1109, 'прохуяшивать'),
    (1110, 'прохуяшиваться'),
    (1107, 'прохуяшить'),
    (1108, 'прохуяшиться'),
    (1111, 'разблядоваться'),
    (1112, 'раздрочить'),
    (1113, 'раздрочиться'),
    (1114, 'раззалупаться'),
    (1115, 'разнохуйственно'),
    (1116, 'разъебать'),
    (1117, 'разъебаться'),
    (1120, 'разъебашивать'),
    (1121, 'разъебашиваться'),
    (1118, 'разъебашить'),
    (1119, 'разъебашиться'),
    (1124, 'разъебенивать'),
    (1125, 'разъебениваться'),
    (1122, 'разъебенить'),
    (1123, 'разъебениться'),
    (1126, 'распиздить'),
    (1127, 'распиздиться'),
    (1128, 'распиздовать'),
    (1129, 'распиздоваться'),
    (1130, 'распиздовывать'),
    (1131, 'распиздовываться'),
    (1140, 'распиздон'),
    (1132, 'распиздохать'),
    (1133, 'распиздохаться'),
    (1134, 'распиздохивать'),
    (1135, 'распиздохиваться'),
    (1138, 'распиздошивать'),
    (1139, 'распиздошиваться'),
    (1136, 'распиздошить'),
    (1137, 'распиздошиться'),
    (1141, 'распиздяй'),
    (1144, 'расхуяривать'),
    (1145, 'расхуяриваться'),
    (1142, 'расхуярить'),
    (1143, 'расхуяриться'),
    (1148, 'расхуячивать'),
    (1149, 'расхуячиваться'),
    (1146, 'расхуячить'),
    (1147, 'расхуячиться'),
    (1150, 'сдрочить'),
    (1151, 'сестроеб'),
    (1152, 'сифилитик'),
    (1153, 'сифилюга'),
    (1154, 'скурвиться'),
    (1155, 'смандить'),
    (1157, 'смандить'),
    (1156, 'смандиться'),
    (1314, 'спам'),
    (1158, 'сперматозавр'),
    (1159, 'спиздеть'),
    (1160, 'стерва'),
    (1161, 'стервоза'),
    (1162, 'сука'),
    (1163, 'суки'),
    (1164, 'сукин'),
    (1165, 'сукины'),
    (1166, 'суходрочка'),
    (1167, 'суходрочкой'),
    (1168, 'сучара'),
    (1169, 'сучий'),
    (1170, 'сучка'),
    (1171, 'сучье'),
    (1172, 'схуякать'),
    (1173, 'схуякаться'),
    (1174, 'схуякивать'),
    (1175, 'схуякиваться'),
    (1178, 'схуяривать'),
    (1179, 'схуяриваться'),
    (1176, 'схуярить'),
    (1177, 'схуяриться'),
    (1182, 'схуячивать'),
    (1180, 'схуячить'),
    (1181, 'схуячиться'),
    (1185, 'съебать'),
    (1186, 'съебаться'),
    (1189, 'съебашивать'),
    (1190, 'съебашиваться'),
    (1187, 'съебашить'),
    (1188, 'съебашиться'),
    (1193, 'съебенивать'),
    (1191, 'съебенить'),
    (1192, 'съебениться'),
    (1183, 'съебывать'),
    (1184, 'съебываться'),
    (1194, 'тварь'),
    (1195, 'толстожопый'),
    (1196, 'толстозадая'),
    (1197, 'торчило'),
    (1198, 'траханье'),
    (1199, 'трахать'),
    (1200, 'трахаться'),
    (1201, 'трахнуть'),
    (1202, 'трахнуться'),
    (1203, 'трепак'),
    (1204, 'триппер'),
    (1208, 'ублюдок'),
    (1209, 'уебать'),
    (1211, 'уебашивать'),
    (1210, 'уебашить'),
    (1212, 'уебенить'),
    (1213, 'уебище'),
    (1205, 'уебывать'),
    (1206, 'уебываться'),
    (1207, 'уебыш'),
    (1214, 'усраться'),
    (1215, 'усрачка'),
    (1216, 'уссать'),
    (1217, 'уссаться'),
    (1218, 'ухуякать'),
    (1219, 'ухуякаться'),
    (1220, 'ухуякивать'),
    (1221, 'ухуякиваться'),
    (1224, 'ухуяривать'),
    (1225, 'ухуяриваться'),
    (1222, 'ухуярить'),
    (1223, 'ухуяриться'),
    (1228, 'ухуячивать'),
    (1229, 'ухуячиваться'),
    (1226, 'ухуячить'),
    (1227, 'ухуячиться'),
    (1232, 'ухуяшивать'),
    (1233, 'ухуяшиваться'),
    (1230, 'ухуяшить'),
    (1231, 'ухуяшиться'),
    (1234, 'фаллос'),
    (1235, 'фекал'),
    (1237, 'фекалии'),
    (1236, 'фекалий'),
    (1238, 'хер'),
    (1239, 'херами'),
    (1240, 'херня'),
    (1242, 'херов'),
    (1241, 'херовина'),
    (1243, 'хрен'),
    (1244, 'хреново'),
    (1245, 'хреновое'),
    (1246, 'хреновый'),
    (1253, 'худоебина'),
    (1254, 'хуебень'),
    (1248, 'хуев'),
    (1255, 'хуев'),
    (1256, 'хуева'),
    (1257, 'хуевато'),
    (1258, 'хуеватый'),
    (1247, 'хуевина'),
    (1249, 'хуево'),
    (1250, 'хуевый'),
    (1259, 'хуеглот'),
    (1260, 'хуегрыз'),
    (1261, 'хуедрыга'),
    (1251, 'хуек'),
    (1262, 'хуемудрие'),
    (1263, 'хуемыслие'),
    (1264, 'хуеньки'),
    (1265, 'хуеплет'),
    (1266, 'хуесос'),
    (1267, 'хуета'),
    (1268, 'хуетень'),
    (1269, 'хуец'),
    (1252, 'хуечек'),
    (1275, 'хуи'),
    (1270, 'хуила'),
    (1276, 'хуило'),
    (1271, 'хуиный'),
    (1272, 'хуистый'),
    (1273, 'хуишко'),
    (1274, 'хуище'),
    (1278, 'хуй'),
    (1277, 'хуйло'),
    (1311, 'хуйло'),
    (1280, 'хуйнуть'),
    (1282, 'хуйню'),
    (1281, 'хуйня'),
    (1279, 'хуйство'),
    (1283, 'хули'),
    (1286, 'хуюживать'),
    (1287, 'хуюживаться'),
    (1284, 'хуюжить'),
    (1285, 'хуюжиться'),
    (1288, 'хуюшки'),
    (1289, 'хуя'),
    (1290, 'хуяк'),
    (1291, 'хуякать'),
    (1292, 'хуями'),
    (1293, 'хуярить'),
    (1294, 'хуяриться'),
    (1295, 'хуястый'),
    (1296, 'хуячий'),
    (1297, 'хуячить'),
    (1298, 'хуячиться'),
    (1299, 'хуяшить'),
    (1300, 'целка'),
    (1301, 'целку'),
    (1302, 'целочка'),
    (1303, 'черножопые'),
    (1304, 'чернозадый'),
    (1305, 'член'),
    (1306, 'шалава'),
    (1307, 'шлюха'),
    (1308, 'шмара'),
    (1309, 'шмарить'),
    (1310, 'шмариться');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "ban";
DROP TABLE "chat";
DROP TABLE "user";
-- +goose StatementEnd
